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
	database, err := f._initDatabase()
	if err != nil {
		log.Print("Failed to load database")
		return false, err
	}

	log.Printf("Saving command %s", feedback)

	statement, err := database.Prepare("INSERT INTO feedback (from_user, to_user, feedback) VALUES (?, ?, ?)")
	if err != nil {
		log.Print("Error preparing insert query")
		return false, err
	}

	_, err = statement.Exec(fromUserId, toUserId, feedback)
	if err != nil {
		log.Print("Failed to insert the feedback into table")
		return false, err
	}

	return true, nil
}

func (f feedbackDatabaseBackend) _initDatabase() (*sql.DB, error) {
	log.Print("Preparing database...")

	database, err := sql.Open("sqlite3", "./feedbacks.db")

	if err != nil {
		log.Print("Error while opening database")
		return nil, err
	}

	statement, err := database.Prepare(`
		CREATE TABLE IF NOT EXISTS feedback (
			id INTEGER PRIMARY KEY,
			from_user TEXT,
			to_user TEXT,
			feedback TEXT
		)`,
	)

	if err != nil {
		log.Print("Error while creating feedback table")
		return nil, err
	}

	statement.Exec()
	return database, nil
}
