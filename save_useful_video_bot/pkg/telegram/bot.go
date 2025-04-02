package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"save_useful_video_bot/configs"
	"save_useful_video_bot/pkg/scheduler"
	"save_useful_video_bot/pkg/storage"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	storage *storage.Storage
	config  *configs.Config
}

func NewBot(bot *tgbotapi.BotAPI, db *storage.Storage, cfg *configs.Config) *Bot {
	return &Bot{bot: bot, storage: db, config: cfg}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	scheduler.RefreshJob(b.bot, b.storage)

	updates, err := b.initUpdatesChannel()
	if err != nil {
		return err
	}

	b.handleUpdate(updates)

	return nil
}
func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u), nil
}
