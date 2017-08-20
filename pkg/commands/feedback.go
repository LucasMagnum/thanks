package commands

import (
    "errors"

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

func (f *FeedbackCommand) Process(form form) (string, error) {
    users := f.interactor.GetUsersFromText(form.Get("text"))
    user := entities.User{
        Id: form.Get("user_id"),
        Name: form.Get("user_name"),
    }

    if _, err := f.interactor.IsValidUsers(user, users); err != nil {
        switch err {
        case interactors.ErrUsersNotFound, interactors.ErrSelfFeedback:
            return err.Error(), nil
        default:
            return "", errors.New("Unknown error")
        }
    }

    // TODO: Insert register into database
    return "Congrats! You feedback was registered with success", nil
}


func NewFeedbackCommand() FeedbackCommand {
    return FeedbackCommand{
        interactor: &interactors.FeedbackInteractor{},
    }
}
