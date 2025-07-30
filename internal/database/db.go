package database

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/yourusername/myparser/internal/parsing"
	"github.com/yourusername/myparser/internal/structs"
	"log"
	"time"
)

var DB *sql.DB

func Init(dbURL string) error {
	var err error
	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return DB.PingContext(ctx)
}

func UpdateLastReview(id int64, placeLink, newAuthor, newText string, newStars int) error {
	_, err := DB.Exec(
		`UPDATE links
			   SET last_author = $1, last_text = $2, last_stars  = $3
			   WHERE user_id = $4 AND place_url = $5`, newAuthor, newText, newStars, id, placeLink,
	)
	if err != nil {
		return err
	}
	return nil
}

func NewLink(id int64, link string) error {
	err, lastAuthor, lastText, lastStars := parsing.ParsingFirstReview(link)
	if err != nil {
		return err
	}
	_, err1 := DB.Exec(
		`INSERT INTO links (user_id, place_url, last_author, last_text, last_stars) VALUES ($1, $2, $3, $4, $5)`, id, link, lastAuthor, lastText, lastStars,
	)
	if err1 != nil {
		return err1
	}
	return nil
}

func ExistsLink(id int64, link string) (error, bool) {
	var exists bool
	err := DB.QueryRow(`
        SELECT EXISTS (
            SELECT 1 FROM links 
            WHERE user_id = $1 AND place_url = $2
        )
    `, id, link).Scan(&exists)
	if err != nil {
		return err, false
	}
	return nil, exists
}

func AllUserLinks(id int64) []string {
	var placeURLs []string
	rows, err := DB.Query("SELECT place_url FROM links WHERE user_id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var url string
		if err1 := rows.Scan(&url); err1 != nil {
			log.Fatal(err1)
		}
		placeURLs = append(placeURLs, url)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return placeURLs
}

func AllLinks() (error, []*structs.Link) {
	var links []*structs.Link
	rows, err := DB.Query("SELECT * FROM links")
	if err != nil {
		return err, links
	}
	defer rows.Close()
	for rows.Next() {
		link := structs.NewLink()
		if err1 := rows.Scan(&link.UserID, &link.Link, &link.LastAuthor, &link.LastText, &link.LastStars); err1 != nil {
			log.Fatal(err1)
		}
		links = append(links, link)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return nil, links
}
