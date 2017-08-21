package app

import (
	"bytes"
	"fmt"
)

type FeedbackHandler struct {
	interactor feedbackInteractor
}

func (f *FeedbackHandler) ProcessCommand(command command) (string, error) {
	if err := f.interactor.validateCommand(command); err != nil {
		return "", err
	}

	users := f.interactor.parseUsersFromText(command.text)

	//TODO: Save user into database

	responseText := f.generateSuccessMessage(users)
	return responseText, nil
}

func (f *FeedbackHandler) generateSuccessMessage(users []user) string {
	var message bytes.Buffer

	for _, user := range users {
		message.WriteString(
			fmt.Sprintf("Congratulations @%s! You earned +1 point!\n", user.name),
		)
	}

	return message.String()
}

func NewFeedbackHandler() FeedbackHandler {
	return FeedbackHandler{
		interactor: feedbackInteractor{},
	}
}
