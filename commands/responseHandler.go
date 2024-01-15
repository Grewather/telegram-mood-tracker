package commands

import (
	"database/sql"
	"fmt"
	"strings"
	"telegram-mood-tracker/database"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ResponseHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update, db *sql.DB) {
	mood := [2]string{"good", "bad"}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	feeling := strings.TrimSpace(update.Message.Text[len("I Feel"):])
	currentTime := time.Now()
	day := currentTime.Weekday()
	if isInArray(mood, feeling) {
		if userDataList := database.GetSthFromDb(day, update.Message.Chat.ID, db); len(userDataList) == 0 {
			msg.Text = fmt.Sprintf("You feel %s", feeling)
			bot.Send(msg)
			database.AddToDb(feeling, update.Message.Chat.ID, db)
		} else {
			msg.Text = fmt.Sprintf("You already told me how you feel today")
			bot.Send(msg)
		}
	} else {
		msg.Text = "I don't know that feeling"
		bot.Send(msg)
	}
}

func isInArray(arr [2]string, target string) bool {
	for _, value := range arr {
		if value == target {
			return true
		}
	}
	return false
}
