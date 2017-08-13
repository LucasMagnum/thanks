package configs

import (
    "log"

    "github.com/kelseyhightower/envconfig"
)

// APIConfig control the API basic behaviours
type APIConfig struct {
    AllowedTeamDomain string
    CheckDomain bool
    Port int
}

// CommandsConfig control the Commands behaviours, as
// FeedbackCommand name and RankingCommand name and the
// default ResponseType for a command
type CommandsConfig struct {
    FeedbackCommand string
    RankingCommand string
    ResponseType string
}

// MessagesConfig control the returned messages for the commands
type MessagesConfig struct {
    HandlerNotFound string
    SelfFeedback string
    SuccessFeedback string
    UsersNotFound string
}

var API APIConfig
var Commands CommandsConfig
var Messages MessagesConfig

func init(){
    log.Print("Loading configs")

    err := envconfig.Process("thanksapi", &API)

    if err != nil {
        log.Fatalln("Failed to load API configs")
    }

    err = envconfig.Process("thanksapi", &Commands)
    if err != nil {
        log.Fatalln("Failed to load Commands configs")
    }

    err = envconfig.Process("thanksapi", &Messages)
    if err != nil {
        log.Fatalln("Failed to load Messages configs")
    }

    log.Print("Loaded configs for API, Commands and Messages")
}

