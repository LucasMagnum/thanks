package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lucasmagnum/thanks-api/commands"
    "github.com/lucasmagnum/thanks-api/configs"
	"github.com/lucasmagnum/thanks-api/handlers"
)


func main() {
	fmt.Println("Starting listening 0.0.0.0:4390")
	http.HandleFunc("/", HandleSlackCommand)
	http.ListenAndServe(":4390", nil)
}

func HandleSlackCommand(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if !configs.DomainAllowed(r.Form.Get("team_domain")) {
		http.Error(w, "Domain requests not allowed", 403)
	}

	requestUser := handlers.RequestUser{
		UserID:   r.Form.Get("user_id"),
		Username: r.Form.Get("user_name"),
	}
	commandName := r.Form.Get("command")
	commandText := r.Form.Get("text")

	commandHandler := handlers.Get(commandName)
	responseTxt := commandHandler.Process(commandText, requestUser)

	response := map[string]string{
		"text":          responseTxt.Content,
		"response_type": "in_channel",
	}

	jsonResponse, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
