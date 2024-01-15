package database

import (
	"database/sql"
	"log"
	"strconv"
	"time"
)

func InitDb() (*sql.DB, error) {
	now := time.Now()
	year, week := now.ISOWeek()

	db, err := sql.Open("sqlite3", strconv.Itoa(year)+strconv.Itoa(week)+".db")
	if err != nil {
		log.Fatal(err)
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS moods (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		chatid int,
		day TEXT,
		mood TEXT
	);`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
