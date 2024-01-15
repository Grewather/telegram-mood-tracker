// addSthToDb.go

package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func AddToDb(feeling string, chatid int64, db *sql.DB) {
	today := time.Now()
	weekDayName := today.Weekday().String()

	insertDataQuery := "INSERT INTO moods (chatid, day, mood) VALUES (?, ?, ?);"
	_, err := db.Exec(insertDataQuery, chatid, weekDayName, feeling)
	if err != nil {
		log.Fatal(err)
	}
}
