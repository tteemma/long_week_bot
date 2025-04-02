package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"save_useful_video_bot/pkg/scheduler"
	"save_useful_video_bot/pkg/storage"
	"time"
)

const (
	commandStart  = "start"
	commandChange = "change"
)

func (b *Bot) handleUpdate(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue
		}
		b.handleMessage(update.Message)
	}
}
func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	chatID := message.Chat.ID
	state, err := b.storage.GetUser(chatID)

	if err != nil {
		log.Printf("[%s] %s", message.From.UserName, err.Error())
		return
	}

	if state == nil {
		state = &storage.UserState{}
		msg := tgbotapi.NewMessage(message.Chat.ID, b.config.Messages.Start)
		_, err := b.bot.Send(msg)
		if err != nil {
			log.Println(err)
		}
		return
	}

	if state.Name == "" {
		state.Name = message.Text
		state.IsWaitDate = true
		err := b.storage.SaveUser(chatID, state)
		if err != nil {
			log.Println(err)
			return
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(b.config.Messages.AskDOB, message.Text))
		_, err = b.bot.Send(msg)
		if err != nil {
			log.Println(err)
		}
		return
	}

	if state.IsWaitDate {
		birthDate, err := time.Parse("2006-01-02", message.Text)

		if err != nil {
			response := fmt.Sprintf(b.config.Messages.InvalidDOBFormat, state.Name)
			msg := tgbotapi.NewMessage(message.Chat.ID, response)
			_, err = b.bot.Send(msg)
			if err != nil {
				log.Println(err)
			}
			return
		}
		if (birthDate.Year() <= (time.Now().Year() - 120)) || (birthDate.After(time.Now())) {
			msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(b.config.Messages.InvalidDOB, state.Name))
			msg.ReplyToMessageID = message.MessageID
			_, err := b.bot.Send(msg)
			if err != nil {
				log.Println(err)
			}
			return
		}

		state.BirthDate = birthDate
		state.IsWaitDate = false
		err = b.storage.SaveUser(chatID, state)
		if err != nil {
			log.Println(err)
			return
		}

		weekLived := calculateWeeksLived(birthDate)
		response := fmt.Sprintf(b.config.Messages.WeeksLived, state.Name, weekLived)
		msg := tgbotapi.NewMessage(message.Chat.ID, response)
		_, err = b.bot.Send(msg)
		if err != nil {
			log.Println(err)
		}

		if state.CronID != 0 {
			scheduler.RemoveJob(chatID)
		}

		state.CronID = scheduler.ScheduleWeekly(b.bot, message.Chat.ID, state.BirthDate, state.Name)
		b.storage.SaveUser(chatID, state)
		return
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, b.config.Messages.DOBAlreadySet)
	_, err = b.bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
	return
}

func (b *Bot) handleCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.config.Messages.UnknownCommand)

	switch message.Command() {
	case commandStart:
		msg.Text = b.config.Messages.Start
		state := &storage.UserState{}
		err := b.storage.SaveUser(message.Chat.ID, state)
		if err != nil {
			log.Println(err)
			return
		}
		_, err = b.bot.Send(msg)
		if err != nil {
			log.Println(err)
		}
	case commandChange:
		state, err := b.storage.GetUser(message.Chat.ID)
		if err != nil {
			log.Println(err)
			return
		}
		if state != nil {
			if state.CronID != 0 {
				scheduler.RemoveJob(message.Chat.ID)
			}
			state.BirthDate = time.Time{}
			state.IsWaitDate = true
			err = b.storage.SaveUser(message.Chat.ID, state)
			if err != nil {
				log.Println(err)
				return
			}
			msg.Text = fmt.Sprintf(b.config.Messages.ChangeDOB, state.Name)
			_, err = b.bot.Send(msg)
			if err != nil {
				log.Println(err)
			}
			return
		}
		msg.Text = b.config.Messages.StartWith
		_, err = b.bot.Send(msg)
		if err != nil {
			log.Println(err)
		}

	default:
		_, err := b.bot.Send(msg)
		if err != nil {
			log.Println(err)
		}
	}
}
func calculateWeeksLived(dateStr time.Time) int {
	duration := time.Since(dateStr)
	weeks := int(duration.Hours() / 24 / 7)

	return weeks
}
