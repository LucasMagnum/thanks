package app

import (
	"bytes"
	"fmt"
)

type feedbackHandler struct {
	interactor feedbackInteractor
	backend    feedbackBackend
}

func (f *feedbackHandler) ProcessCommand(command command) (string, error) {
	if err := f.interactor.validateCommand(command); err != nil {
		return "", err
	}

	users := f.interactor.parseUsersFromText(command.text)

	// Save into database
	f.backend.save(command)

	responseText := f.generateSuccessMessage(users)
	return responseText, nil
}

func (f *feedbackHandler) generateSuccessMessage(users []user) string {
	var message bytes.Buffer

	for _, user := range users {
		message.WriteString(
			fmt.Sprintf("Congratulations @%s! You earned +1 point!\n", user.name),
		)
	}

	return message.String()
}

func NewFeedbackHandler() feedbackHandler {
	return feedbackHandler{
		interactor: feedbackInteractor{},
		backend:    feedbackDatabaseBackend{},
	}
}
