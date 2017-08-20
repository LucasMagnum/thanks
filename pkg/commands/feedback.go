package commands

import (
	"bytes"
	"fmt"

	"github.com/lucasmagnum/thanks/pkg/entities"
	"github.com/lucasmagnum/thanks/pkg/interactors"
)

type form interface {
	Get(key string) string
}

type interactor interface {
	GetUsersFromText(text string) []entities.User
	ValidateUsers(user entities.User, users []entities.User) error
}

type FeedbackCommand struct {
	interactor interactor
}

func (f *FeedbackCommand) Process(form form) (Response, error) {
	users := f.interactor.GetUsersFromText(form.Get("text"))
	user := entities.User{
		Id:   form.Get("user_id"),
		Name: form.Get("user_name"),
	}

	if err := f.interactor.ValidateUsers(user, users); err != nil {
		return Response{
			Text:         err.Error(),
			ResponseType: "ephemeral",
		}, nil
	}

	responseText := generateSuccessMessage(users)

	// TODO: Insert register into database
	return Response{Text: responseText}, nil
}

func NewFeedbackCommand() FeedbackCommand {
	return FeedbackCommand{
		interactor: &interactors.FeedbackInteractor{},
	}
}

func generateSuccessMessage(users []entities.User) string {
	var message bytes.Buffer

	for _, user := range users {
		message.WriteString(
			fmt.Sprintf("Congratulations @%s! You earned +1 point!\n", user.Name),
		)
	}

	return message.String()
}
