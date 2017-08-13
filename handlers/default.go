package handlers

import (
    "github.com/lucasmagnum/thanks-api/configs"
)


// emptyHandler is used as default handler for unknown requests
type emptyHandler struct{}


// Process empty handler should return the HandlerNotFound message
func (d emptyHandler) Process(commandText string, slackUser SlackUser) Result {
	return Result{
		Content: configs.Messages.HandlerNotFound,
	}
}
