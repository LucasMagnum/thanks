package commands

import (
	"bytes"
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
        return usersNotFoundMessage()
    }

	for _, user := range users {
        if user.UserId == slackUser.UserId {
            // When the user is trying to send a feedback
            // to himself, we should return the self feedback message
            if usersCount == 1 {
                return selfFeedbackMessage()
            }
            continue
        }

        successMessage := successFeedbackMessage(user.Username)
        resultContent.WriteString(successMessage.Content)
	}

    return message(resultContent.String())
}

// parseUsers receives the command text and return the a slice
// of SlackUser
func parseUsers(commandText string) (users []handlers.SlackUser) {
    for _, user := range getUsersFromText(commandText) {
        userId, username := parseUserData(user)

        users = append(users, handlers.SlackUser{
            UserId:   userId,
            Username: username,
        })
    }

    return
}

// getUsersFromText get user array from commandText
// the commandText could contain one or more users
func getUsersFromText(commandText string) ([]string){
    userRegex := regexp.MustCompile("<?@([a-zA-Z0-9-_.|])+>?")
    return userRegex.FindAllString(commandText, -1)
}


// parseUserData parse data and return the userId and username
// this function expect the username in the format <@ID|Name> or
// <@Username>
func parseUserData(userData string) (userId, username string) {
    clearRegex := regexp.MustCompile("[<@>]")
    cleanedUsername := clearRegex.ReplaceAllString(userData, "")

    userId, username = cleanedUsername, cleanedUsername

    // User has the escaped value <@userid|username>
    if strings.Contains(cleanedUsername, "|") {
        splitUser := strings.Split(cleanedUsername, "|")

        userId = splitUser[0]
        username = splitUser[1]
    }

    return
}

