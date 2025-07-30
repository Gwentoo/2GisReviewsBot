package bot

import (
	"github.com/yourusername/myparser/internal/app/handlers"
	tele "gopkg.in/telebot.v3"
	"log"
	"time"
)

var commands = []tele.Command{
	{
		Text:        "info",
		Description: "Информация",
	},
	{
		Text:        "new",
		Description: "Добавить новое место",
	},
	{
		Text:        "places",
		Description: "Добавленные места",
	},
	{
		Text:        "del",
		Description: "Удалить место",
	},
}

func NewBot(token string) (*tele.Bot, error) {

	bot, err := tele.NewBot(tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Println(err)
	}
	err2 := bot.SetCommands(commands)
	if err2 != nil {
		log.Println(err2)
	}

	handlers.RegisterTextHandlers(bot)
	handlers.RegisterCommonHandlers(bot)

	return bot, nil
}
