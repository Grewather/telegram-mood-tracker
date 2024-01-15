package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"telegram-scheduler/commands"
	"telegram-scheduler/database"
	"telegram-scheduler/reminders"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	godotenv.Load()
	token := os.Getenv("TOKEN")
	db, err := database.InitDb()
	if err != nil {
		log.Fatal(err)
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false // set it to true to turn on telegram bot debug mode

	log.Printf("Authorized on account %s", bot.Self.UserName)

	go reminders.GetWeeklySummary(db, bot)

	go reminders.StartCron(bot)

	commands.HandleCommands(bot, db)
	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
