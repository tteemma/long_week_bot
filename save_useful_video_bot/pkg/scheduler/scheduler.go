package scheduler

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/robfig/cron/v3"
	"log"
	"save_useful_video_bot/pkg/storage"
	"time"
)

var c = cron.New(cron.WithSeconds())
var jobs = make(map[int64]cron.EntryID)

func ScheduleWeekly(b *tgbotapi.BotAPI, chatId int64, birthDate time.Time, name string) cron.EntryID {

	RemoveJob(chatId)

	//jobID, err := c.AddFunc("*/1 * * * * *", func() {
	jobID, err := c.AddFunc("0 0 7 * * MON", func() {
		weeksLived := int(time.Since(birthDate).Hours() / (24 * 7))
		msg := tgbotapi.NewMessage(chatId, fmt.Sprintf("%s, вы прожили %d недель!", name, weeksLived))
		_, err := b.Send(msg)
		if err != nil {
			log.Println(err)
		}
	})
	if err != nil {
		log.Printf("Ошибка добавления задачи в cron: %v", err)
		return 0
	}

	jobs[chatId] = jobID
	log.Printf("Запланировано обновление для %d", chatId)

	c.Start()

	return jobID
}

func RemoveJob(chatId int64) {
	if jobId, exist := jobs[chatId]; exist {
		c.Remove(jobId)
		delete(jobs, chatId)
		log.Printf("Удалена старая дата рождения для %d", chatId)
	}
}
func RefreshJob(b *tgbotapi.BotAPI, storage *storage.Storage) {
	users, err := storage.GetAllUsersAfterRefresh()
	if err != nil {
		log.Println(err)
		return
	}

	for chatID, user := range users {
		if user.CronID != 0 && !user.BirthDate.IsZero() {
			RemoveJob(chatID)

			user.CronID = ScheduleWeekly(b, chatID, user.BirthDate, user.Name)
			err := storage.SaveUser(chatID, user)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
