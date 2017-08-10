package handlers

type emptyHandler struct{}

func (d emptyHandler) Process(commandText string, requestUser RequestUser) Result {
    return Result{
        Content: "Unfortunately we can't handle your request =(",
    }
}
