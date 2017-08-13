package tests

import (
    "fmt"
    "reflect"
    "testing"

    "github.com/lucasmagnum/thanks-api/configs"
    "github.com/lucasmagnum/thanks-api/commands"
    "github.com/lucasmagnum/thanks-api/handlers"
)


func TestFeedbackCommand(t *testing.T){

    // Should register the command on init
    feedbackCommand := configs.Commands.FeedbackCommand
    handler := handlers.Get(feedbackCommand)

    if reflect.TypeOf(handler) != reflect.TypeOf(commands.FeedbackCommand{}) {
        t.Error("handler registered isn't the FeedbackCommand")
    }

    slackUser := handlers.SlackUser{
        UserId: "123",
        Username: "test-user",
    }

    // Should return UsersNotFound message when the text doesn't contains users
    commandText := "This text doesn't contains users"
    result := handler.Process(commandText, slackUser)

    notFoundMessage := fmt.Sprintf(configs.Messages.UsersNotFound, feedbackCommand)
    if result.Content != notFoundMessage {
        t.Error("Expected UsersNotFound message, got", result.Content)
    }

    // Should SelfFeedback message when user try to feedback himself
    commandText = "<@123|test-user> for the help."
    result = handler.Process(commandText, slackUser)

    selfFeedbackMessage := configs.Messages.SelfFeedback
    if result.Content != selfFeedbackMessage {
        t.Error("Expected SelfFeedback message, got", result.Content)
    }

    // Should return success message when user is valid
    commandText = "<@124|test.user.2> for the help."
    result = handler.Process(commandText, slackUser)

    successMessage := fmt.Sprintf(configs.Messages.SuccessFeedback, "test.user.2")
    if result.Content != successMessage {
        t.Error("Expected SuccessFeedback message, got", result.Content)
    }

}


