package interactors

import (
    "errors"

    "github.com/lucasmagnum/thanks/pkg/entities"
)

var ErrUsersNotFound = errors.New("Users not found")
var ErrSelfFeedback = errors.New("Self feedback not allowed")


type FeedbackInteractor struct {}

func (f *FeedbackInteractor) GetUsersFromText(text string) []entities.User {
    return []entities.User{}
}

func (f *FeedbackInteractor) IsValidUsers(user entities.User, users []entities.User) (bool, error) {
    if len(users) == 0 {
        return false, ErrUsersNotFound
    }

    if findUser(user, users) {
        return false, ErrSelfFeedback
    }

    return true, nil
}

func findUser(user entities.User, users []entities.User) bool {
    for _, userSlice := range users {
        if userSlice.Id == user.Id {
            return true
        }
    }
    return false
}

