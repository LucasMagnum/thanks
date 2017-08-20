package api

import (
    "encoding/json"
    "net/http"

    "github.com/lucasmagnum/thanks/pkg/commands"
)

func FeedbackHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()

    command := commands.NewFeedbackCommand()
    response, err := command.Process(r.Form)

    if err != nil {
        http.Error(w, err.Error(), 500)
    }

    jsonResponse, _ := json.Marshal(map[string]string{
        "text": response,
        "response_type": "in_channel",
    })

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)
}
