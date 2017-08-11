package commands

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/lucasmagnum/thanks-api/configs"
	"github.com/lucasmagnum/thanks-api/handlers"
)

var feedbackCommandName = configs.Commands.FeedbackCommand


func init() {
    var command FeedbackCommand
    handlers.Register(feedbackCommandName, command)
}


type FeedbackCommand struct{}

func (f FeedbackCommand) Process(commandText string, slackUser handlers.SlackUser) handlers.Result {
	var resultContent bytes.Buffer

	users := parseUsers(commandText)
    usersCount := len(users)

    if usersCount == 0 {
        return handlers.Result{
            Content: fmt.Sprintf(configs.Messages.UsersNotFound, feedbackCommandName),
        }
    }

	for _, user := range users {
        if user.UserId == slackUser.UserId {

            // When the user is trying to send a feedback
            // to himself, we should return the self feedback message
            if usersCount == 1 {
                return handlers.Result{
                    Content: configs.Messages.SelfFeedback,
                }
            }

            continue
        }

        text := fmt.Sprintf(configs.Messages.SuccessFeedback, user.Username)
        resultContent.WriteString(text)
	}

    return handlers.Result{
        Content: resultContent.String(),
    }

}


func parseUsers(commandText string) (users []handlers.SlackUser) {
	userRegex := regexp.MustCompile("@([a-zA-Z0-9].)+")

	for _, user := range userRegex.FindAllString(commandText, -1) {
		userId, userName := user, user[1:]

		// User has the escaped value <@userid|username>
		if strings.Contains(user, "|") {
            cleanedUsername := strings.Replace(user, ">", "", -1)
			splitUser := strings.Split(cleanedUsername, "|")

			userId = splitUser[0][1:]
			userName = splitUser[1]
		}

		users = append(users, handlers.SlackUser{
			UserId:   userId,
			Username: userName,
		})
	}

	return
}
