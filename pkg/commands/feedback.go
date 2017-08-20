package commands

import (
	"github.com/lucasmagnum/thanks/pkg/entities"
	"github.com/lucasmagnum/thanks/pkg/interactors"
)

type form interface {
	Get(key string) string
}

type interactor interface {
	GetUsersFromText(text string) []entities.User
	IsValidUsers(user entities.User, users []entities.User) (bool, error)
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

	if _, err := f.interactor.IsValidUsers(user, users); err != nil {
		return Response{
            Text: err.Error(),
            ResponseType: "ephemeral",
        }, nil
	}

	// TODO: Insert register into database
	return Response{
        Text: "Congrats! You feedback was registered with success",
    }, nil
}

func NewFeedbackCommand() FeedbackCommand {
	return FeedbackCommand{
		interactor: &interactors.FeedbackInteractor{},
	}
}
