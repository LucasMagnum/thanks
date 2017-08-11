package configs

import (
    "log"

    "github.com/kelseyhightower/envconfig"
)

type APIConfig struct {
    AllowedTeamDomain string
    CheckDomain bool
    Port int
}

type CommandsConfig struct {
    FeedbackCommand string
    RankingCommand string
    ResponseType string
}

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

