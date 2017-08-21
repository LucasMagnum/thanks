package app

import (
	"errors"
	"regexp"
	"strings"
)

var ErrUsersNotFound = errors.New("Users not found")
var ErrSelfFeedback = errors.New("Self feedback not allowed")

type feedbackInteractor struct{}

func (f *feedbackInteractor) validateCommand(command command) error {
	users := f.parseUsersFromText(command.text)

	if len(users) == 0 {
		return ErrUsersNotFound
	}

	commandUser := user{
		id:   command.userId,
		name: command.userName,
	}

	if hasUser(users, commandUser) {
		return ErrSelfFeedback
	}

	return nil
}

func (f *feedbackInteractor) parseUsersFromText(text string) []user {
	userRegex := regexp.MustCompile("<@([\\w-_.|])+>")
	userStrings := userRegex.FindAllString(text, -1)

	usersMap := make(map[string]user)

	for _, userData := range userStrings {
		userId, userName := f.parseUserData(userData)

		// Avoid duplicate results
		if _, ok := usersMap[userId]; !ok {
			usersMap[userId] = user{
				id:   userId,
				name: userName,
			}
		}
	}

	users := []user{}

	for _, user := range usersMap {
		users = append(users, user)
	}

	return users
}

func (f *feedbackInteractor) parseUserData(userData string) (string, string) {
	clearRegex := regexp.MustCompile("[<@>]")
	cleanedData := clearRegex.ReplaceAllString(userData, "")

	splitUser := strings.Split(cleanedData, "|")

	userId := splitUser[0]
	userName := splitUser[1]

	return userId, userName
}

func hasUser(users []user, commandUser user) bool {
	for _, userSlice := range users {
		if userSlice.Equal(commandUser) {
			return true
		}
	}
	return false
}
