package app

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type feedbackBackend interface {
	save(feedback string, fromUserId string, toUserId string) (bool, error)
}

type feedbackDatabaseBackend struct{}

func (f feedbackDatabaseBackend) save(feedback string, fromUserId string, toUserId string) (bool, error) {
	database, _ := sql.Open("sqlite3", "./feedbacks.db")
	statement, _ := database.Prepare(`
	CREATE TABLE IF NOT EXISTS feedback
	(
		id INTEGER PRIMARY KEY,
		from_user TEXT,
		to_user TEXT,
		feedback TEXT
	)`)
	statement.Exec()

	statement, _ = database.Prepare("INSERT INTO feedbacks (from_user, to_user, feedback) VALUES (?, ?, ?)")
	statement.Exec("Text", "Tet", "Text")

	log.Printf("Saving command %s", feedback)
	return true, nil
}
