package period

import (
	"fmt"
	"github.com/yourusername/myparser/internal/database"
	"github.com/yourusername/myparser/internal/parsing"
	tele "gopkg.in/telebot.v3"
	"log"
)

func CheckNewReviews(bot *tele.Bot) error {
	err, links := database.AllLinks()
	if err != nil {
		return err
	}
	for _, link := range links {
		var first bool
		err1, newAuthor, newText, newStars := parsing.ParsingReviews(link.Link, link.LastAuthor, link.LastText, link.LastStars)
		if err1 != nil {
			return err1
		}
		if len(newAuthor) == 0 {
			continue
		}
		for i := 0; i < len(newAuthor); i++ {
			var mes, stars string
			for j := 0; j < newStars[i]; j++ {
				stars += "★"
			}
			for j := 0; j < 5-newStars[i]; j++ {
				stars += "✰"
			}
			mes += "Новый отзыв:\n"
			name, address := parsing.ParsingName(link.Link)
			mes += fmt.Sprintf("Адрес: %s\n", address)
			mes += fmt.Sprintf("Место: %s\n\n", name)
			mes += newAuthor[i] + "\n" + stars + "\n" + newText[i]
			bot.Send(tele.ChatID(link.UserID), mes)

			if !first {
				err2 := database.UpdateLastReview(link.UserID, link.Link, newAuthor[i], newText[i], newStars[i])
				if err2 != nil {
					log.Fatal(err2)
				}
				first = true
			}

		}
	}
	return nil

}
