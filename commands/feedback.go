package commands

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/lucasmagnum/thanks-api/configs"
	"github.com/lucasmagnum/thanks-api/handlers"
)

func init() {
	var command FeedbackCommand
	handlers.Register(configs.FeedbackCommandName, command)
}

type FeedbackCommand struct{}

func (f FeedbackCommand) Process(commandText string, slackUser handlers.SlackUser) handlers.Result {
	var resultContent bytes.Buffer

	users, _ := parseCommand(commandText)

	if len(users) == 0 {
		return handlers.Result{
			Content: fmt.Sprintf(
				`Hej! Could you specify one or more users for your feedback? :D
                Ex: %s @username <feedback>`, configs.FeedbackCommandName),
		}
	}

	for _, user := range users {
		text := fmt.Sprintf("Congratulations @%s! You earned +1 feedback point\n", user.Username)
		resultContent.WriteString(text)
	}

	result := handlers.Result{
		Content: resultContent.String(),
	}

	return result
}

func parseCommand(commandText string) ([]handlers.SlackUser, string) {
	var users []handlers.SlackUser

	userRegex := regexp.MustCompile("@([a-zA-Z0-9].)+")

	for _, user := range userRegex.FindAllString(commandText, -1) {
		userId, userName := user, user[1:]

		// User has the escaped value @userid|username
		if strings.Contains(user, "|") {
			splitUser := strings.Split(user, "|")
			userId = splitUser[0][1:]
			userName = splitUser[1]
		}

		users = append(users, handlers.SlackUser{
			UserId:   userId,
			Username: userName,
		})
	}

	return users, commandText
}
