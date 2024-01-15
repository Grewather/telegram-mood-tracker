package reminders

import (
	"database/sql"
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/robfig/cron/v3"
	"telegram-mood-tracker/database"
)

func GetWeeklySummary(db *sql.DB, bot *tgbotapi.BotAPI) {
	c := cron.New()
	_, err := c.AddFunc("55 23 * * *", func() {
		date := time.Now()
		if date.Weekday() == time.Sunday {
			for _, chatID := range UsersToSend {
				feelUnknown := 0
				feelGood := 0
				feelBad := 0

				for _, weekday := range []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday, time.Sunday} {
					userDataList := database.GetSthFromDb(weekday, chatID, db)
					for _, userData := range userDataList {
						switch userData.Mood {
						case "good":
							feelGood++
						case "bad":
							feelBad++
						default:
							feelUnknown++
						}
					}

					if len(userDataList) == 0 {
						fmt.Printf("No data found for %s and ChatID %d\n", weekday, chatID)
						feelUnknown++
					}
				}

				totalDays := float64(len([]time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday, time.Sunday}))
				percentGood := int(float64(feelGood) / totalDays * 100)
				percentBad := int(float64(feelBad) / totalDays * 100)
				percentUnknown := int(float64(feelUnknown) / totalDays * 100)

				message := tgbotapi.NewMessage(chatID, fmt.Sprintf("Weekly Summary: %d%% time you were feeling good, %d%% time you were feeling bad, %d%% time data is undefined", percentGood, percentBad, percentUnknown))
				bot.Send(message)
			}
		}
	})

	if err != nil {
		fmt.Println("Something went wrong while adding a task to a cron:", err)
		return
	}

	c.Start()

	select {}
}
