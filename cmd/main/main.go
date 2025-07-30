package main

import (
	"fmt"
	"github.com/yourusername/myparser/cmd/period"
	"github.com/yourusername/myparser/internal/app/bot"
	"github.com/yourusername/myparser/internal/config"
	"github.com/yourusername/myparser/internal/database"

	"log"
	"time"
)

func main() {
	cfg := config.Load()
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err := database.Init(connStr); err != nil {
		log.Fatal("DB connection failed:", err)
	}
	defer database.DB.Close()
	log.Println("БД подключена")
	Bot, err := bot.NewBot(cfg.TelegramToken)
	if err != nil {
		log.Println(err)
	}
	log.Println("Бот запущен")
	go period.RunPeriodTask(Bot, 1*time.Minute)
	Bot.Start()

}
