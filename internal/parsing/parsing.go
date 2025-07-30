package parsing

import (
	"fmt"
	"github.com/gocolly/colly"
	"unicode/utf8"
)

func ParsingFirstReview(link string) (error, string, string, int) {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
	)

	var author, text string
	var stars int
	var single bool
	c.OnHTML("div._1k5soqfl", func(e *colly.HTMLElement) {
		if !single {
			author = e.DOM.Find("span._16s5yj36").Text()
			text = e.DOM.Find("a._1msln3t").Text()
			if text == "" {
				text = e.DOM.Find("a._1wlx08h").Text()
			}
			if utf8.RuneCountInString(text) > 50 {
				runes := []rune(text)
				text = string(runes[:50])
			}
			stars = e.DOM.Find("div._1fkin5c").Find("span").Length()
			single = true
		}
	})

	url := fmt.Sprintf("%s/tab/reviews", link)
	err := c.Visit(url)
	if err != nil {
		return err, "", "", 0
	}
	return nil, author, text, stars
}

func ParsingReviews(link, lastAuthor, lastText string, lastStars int) (error, []string, []string, []int) {
	var authorArr, textArr []string
	var starsArr []int
	var newRev bool
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
	)
	c.OnHTML("div._1k5soqfl", func(e *colly.HTMLElement) {
		if !newRev {
			var author, text string
			var stars int
			author = e.DOM.Find("span._16s5yj36").Text()
			text = e.DOM.Find("a._1msln3t").Text()
			if text == "" {
				text = e.DOM.Find("a._1wlx08h").Text()
			}
			if utf8.RuneCountInString(text) > 50 {
				runes := []rune(text)
				text = string(runes[:50])
			}
			stars = e.DOM.Find("div._1fkin5c").Find("span").Length()
			if author != lastAuthor || text != lastText || lastStars != stars {
				authorArr = append(authorArr, author)
				textArr = append(textArr, text)
				starsArr = append(starsArr, stars)
			} else {
				newRev = true
			}
		}
	})
	url := fmt.Sprintf("%s/tab/reviews", link)
	err := c.Visit(url)
	if err != nil {
		return err, authorArr, textArr, starsArr
	}
	return nil, authorArr, textArr, starsArr
}

func ParsingName(link string) (string, string) {
	var placeName string
	var placeAddress string
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
	)
	c.OnHTML("div._121zpzx", func(e *colly.HTMLElement) {
		placeName = e.DOM.Find("h1._1x89xo5").Text()
		placeAddress = e.DOM.Find("div._1g2rw7z span._oqoid").Text()
		if placeAddress == "" {
			placeAddress = e.DOM.Find("span._wrdavn a._2lcm958").Text()
		}
		if placeAddress == "" {
			placeAddress = e.DOM.Find("span._oqoid a._2lcm958").Text()
		}
		if placeName != "" && placeAddress == "" {
			placeAddress = placeName
		}
	})

	err := c.Visit(link)
	if err != nil {
		return "", ""
	}

	return placeName, placeAddress
}
