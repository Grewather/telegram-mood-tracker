package commands

import (
	"database/sql"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	// Import the database package
)

func HandleCommands(bot *tgbotapi.BotAPI, db *sql.DB) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 10
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

			switch update.Message.Command() {
			case "help":
				msg.Text = "use /start to start this. Tell me how you feel by typing 'I feel good' or 'I feel bad'"
				bot.Send(msg)
			case "start":
				StartHandler(bot, update)
			default:
				msg.Text = "I don't know that command"
				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
			}
		}
		if strings.HasPrefix(strings.ToLower(update.Message.Text), "i feel") {
			ResponseHandler(bot, update, db)
		}
	}
}
