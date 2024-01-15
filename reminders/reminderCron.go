package reminders

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/robfig/cron/v3"
)

var UsersToSend []int64

func StartCron(bot *tgbotapi.BotAPI) {

	c := cron.New()

	_, err := c.AddFunc("0 16 * * *", func() {
		for _, chatID := range UsersToSend {
			fmt.Println("Sending message to:", chatID)
			message := tgbotapi.NewMessage(chatID, "How are you feeling?")
			bot.Send(message)
		}
	})
	if err != nil {
		fmt.Println("Failed:", err)
		return
	}

	c.Start()

}
