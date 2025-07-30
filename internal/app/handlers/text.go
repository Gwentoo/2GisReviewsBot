package handlers

import (
	"fmt"
	"github.com/yourusername/myparser/internal/database"
	"github.com/yourusername/myparser/internal/parsing"
	"github.com/yourusername/myparser/internal/utils"
	tele "gopkg.in/telebot.v3"
	"strings"
)

func RegisterTextHandlers(bot *tele.Bot) {
	bot.Handle(tele.OnText, func(c tele.Context) error {
		switch UserState[c.Sender().ID] {
		case "newLink":
			return handleNewLink(c)
		default:
			return nil
		}
	})
}

func handleNewLink(c tele.Context) error {
	link := c.Text()
	UserState[c.Sender().ID] = ""
	if strings.Contains(link, "https://go.") {
		link1, err := utils.ExpandShortURL(link)
		if err != nil {
			return err
		}
		link = link1
	}

	parts := strings.Split(link, "?")
	link = parts[0]
	err, exists := database.ExistsLink(c.Sender().ID, link)
	if err != nil {
		return err
	}
	if exists {
		return c.Send("Это место уже добавлено ...")
	}

	placeName, placeAddress := parsing.ParsingName(link)

	if placeName != "" && placeAddress != "" {
		err1 := database.NewLink(c.Sender().ID, link)
		if err1 != nil {
			return err1
		}
		return c.Send(fmt.Sprintf("Добавлено место: %s\nАдрес: %s", placeName, placeAddress))
	} else {
		return c.Send("Некорректная ссылка")
	}

}
