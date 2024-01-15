// database.go

package database

import (
	"database/sql"
	"log"
	"time"
)

type UserData struct {
	ID     int
	ChatID int64
	Day    string
	Mood   string
}

func GetSthFromDb(day time.Weekday, chatID int64, db *sql.DB) []UserData {
	query := "SELECT * FROM moods WHERE day = ? AND chatid = ?;"
	rows, err := db.Query(query, day.String(), chatID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var userDataList []UserData

	for rows.Next() {
		var userData UserData
		err := rows.Scan(&userData.ID, &userData.ChatID, &userData.Day, &userData.Mood)
		if err != nil {
			log.Fatal(err)
		}
		userDataList = append(userDataList, userData)
	}

	return userDataList
}
