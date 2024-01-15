package commands

import (
	"telegram-mood-tracker/reminders"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	reminders.UsersToSend = append(reminders.UsersToSend, update.Message.Chat.ID)
	startMessage := tgbotapi.NewMessage(update.Message.Chat.ID, "Started sending reminders about mood")
	bot.Send(startMessage)
}
