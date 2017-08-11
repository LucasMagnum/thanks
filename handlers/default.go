package handlers

type emptyHandler struct{}

func (d emptyHandler) Process(commandText string, slackUser SlackUser) Result {
	return Result{
		Content: "Unfortunately we can't handle your request =(",
	}
}
