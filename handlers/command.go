package handlers

// User that made the command request
type RequestUser struct {
	UserID   string
	Username string
}

type Result struct {
	Content string
}

// Command interface
type Command interface {
	Process(commandText string, requestUser RequestUser) Result
}
