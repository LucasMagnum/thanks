package handlers

// User that made the command request
type SlackUser struct {
	UserId   string
	Username string
}

type Result struct {
	Content string
}

// Command interface
type Command interface {
	Process(commandText string, slackUser SlackUser) Result
}
