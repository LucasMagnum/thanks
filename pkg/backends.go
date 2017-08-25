package app

import (
	"log"
)

type feedbackBackend interface {
	save(command command) (bool, error)
}

type feedbackDatabaseBackend struct{}

func (f feedbackDatabaseBackend) save(command command) (bool, error) {
	log.Printf("Saving command %s", command.text)
	return true, nil
}
