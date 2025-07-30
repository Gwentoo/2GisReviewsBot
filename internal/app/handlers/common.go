package handlers

import (
	"github.com/yourusername/myparser/internal/database"
	"github.com/yourusername/myparser/internal/parsing"
	tele "gopkg.in/telebot.v3"
)

var UserState = make(map[int64]string)

func RegisterCommonHandlers(bot *tele.Bot) {
	bot.Handle("/info", handleInfo)
	bot.Handle("/new", handleNew)
	bot.Handle("/places", handlePlaces)
	bot.Handle("/start", handleStart)

}

func handleStart(c tele.Context) error {
	UserState[c.Sender().ID] = ""
	return c.Send("Данный бот позволяет отслеживать появление новых отзывов на картах 2ГИС.\nДля добавления места пропишите /new")
}

func handleInfo(c tele.Context) error {
	return c.Send("Бот отправляет сообщения, в котором будут новые отзывы на выбранных местах.\nДля добавления места пропишите /new")
}

func handleNew(c tele.Context) error {
	UserState[c.Sender().ID] = "newLink"
	return c.Send("Отправьте ссылку на место в 2гис\nПример: https://2gis.ru/spb/geo/70030076748037405")
}

func handlePlaces(c tele.Context) error {
	links := database.AllUserLinks(c.Sender().ID)
	if len(links) == 0 {
		return c.Send("Ещё нет добавленных мест. Вы можете добавить места при  помощи /new")
	}
	mes := "Добавленные места:\n\n"

	for _, link := range links {
		name, address := parsing.ParsingName(link)
		mes += "Название: " + name + "\n"
		mes += "Адрес: " + address + "\n\n"
	}
	return c.Send(mes)
}
