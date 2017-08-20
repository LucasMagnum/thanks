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

	if len(response.ResponseType) == 0 {
		response.ResponseType = "in_channel"
	}

	jsonResponse, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
