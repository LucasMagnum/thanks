package interactors

import (
	"errors"
	"regexp"
	"strings"

	"github.com/lucasmagnum/thanks/pkg/entities"
)

var ErrUsersNotFound = errors.New("Users not found")
var ErrSelfFeedback = errors.New("Self feedback not allowed")

type FeedbackInteractor struct{}

func (f *FeedbackInteractor) GetUsersFromText(text string) []entities.User {
	userRegex := regexp.MustCompile("<@([\\w-_.|])+>")
	userStrings := userRegex.FindAllString(text, -1)

	usersMap := make(map[string]entities.User)

	for _, userData := range userStrings {
		userId, userName := parseUserData(userData)

		// Avoid duplicate results
		if _, ok := usersMap[userId]; !ok {
			usersMap[userId] = entities.User{
				Id:   userId,
				Name: userName,
			}
		}
	}

	users := []entities.User{}

	for _, user := range usersMap {
		users = append(users, user)
	}

	return users
}

func (f *FeedbackInteractor) ValidateUsers(user entities.User, users []entities.User) error {
	if len(users) == 0 {
		return ErrUsersNotFound
	}

	if findUser(user, users) {
		return ErrSelfFeedback
	}

	return nil
}

func findUser(user entities.User, users []entities.User) bool {
	for _, userSlice := range users {
		if userSlice.Id == user.Id {
			return true
		}
	}
	return false
}

func parseUserData(userData string) (userId, userName string) {
	clearRegex := regexp.MustCompile("[<@>]")
	cleanedData := clearRegex.ReplaceAllString(userData, "")

	splitUser := strings.Split(cleanedData, "|")

	userId = splitUser[0]
	userName = splitUser[1]

	return
}
