package commands

import (
    "github.com/lucasmagnum/thanks-api/configs"
    "github.com/lucasmagnum/thanks-api/handlers"
)

var rankingCommandName = configs.Commands.RankingCommand


func init() {
    var command RankingCommand
    handlers.Register(rankingCommandName, command)
}


type RankingCommand struct{}

func (r RankingCommand) Process(commandText string, slackUser handlers.SlackUser) handlers.Result {
    return handlers.Result{
        Content: "A.W.E.S.O.M.E",
    }
}

