package commands

import (
    "fmt"

    "github.com/lucasmagnum/thanks-api/configs"
    "github.com/lucasmagnum/thanks-api/handlers"
)


func usersNotFoundMessage() handlers.Result {
    feedbackCommand := configs.Commands.FeedbackCommand
    return message(fmt.Sprintf(configs.Messages.UsersNotFound, feedbackCommand))
}


func selfFeedbackMessage() handlers.Result {
    return message(configs.Messages.SelfFeedback)
}

func successFeedbackMessage(username string) handlers.Result {
    return message(fmt.Sprintf(configs.Messages.SuccessFeedback, username))
}


func message(text string) handlers.Result {
    return handlers.Result {
        Content: text,
    }
}
