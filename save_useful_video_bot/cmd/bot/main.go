package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.etcd.io/bbolt"
	"log"
	"save_useful_video_bot/configs"
	"save_useful_video_bot/pkg/storage"
	"save_useful_video_bot/pkg/telegram"
)

func main() {

	cfg := configs.LoadConfig()

	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	db, err := bbolt.Open("users.db", 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	storage := storage.NewStorage(db)

	telegramBot := telegram.NewBot(bot, storage, cfg)
	if err := telegramBot.Start(); err != nil {
		log.Panic(err)
	}

}
