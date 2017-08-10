package commands

import (
	"fmt"
	"github.com/lucasmagnum/thanks-api/handlers"
)

func init() {
	var command FeedbackCommand
	handlers.Register("/thanks", command)
}

type FeedbackCommand struct{}

func (f FeedbackCommand) Process(commandText string, requestUser handlers.RequestUser) (handlers.Result, error) {
	content := fmt.Sprintf("Congratulations @%s! You called this command!", requestUser.Username)

	result := handlers.Result{
		Content: content,
	}

	return result, nil
}
