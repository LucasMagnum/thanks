package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lucasmagnum/thanks-api/commands"
	"github.com/lucasmagnum/thanks-api/configs"
	"github.com/lucasmagnum/thanks-api/handlers"
)

func main() {
	port := fmt.Sprintf(":%d", configs.API.Port)

	log.Printf("Starting listening 0.0.0.0%s", port)
	http.HandleFunc("/", HandleSlackCommand)
	http.ListenAndServe(port, nil)
}

func HandleSlackCommand(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if configs.API.CheckDomain && configs.API.AllowedTeamDomain != r.Form.Get("team_domain") {
		http.Error(w, "Domain requests not allowed", 403)
	}

	slackUser := handlers.SlackUser{
		UserId:   r.Form.Get("user_id"),
		Username: r.Form.Get("user_name"),
	}

	commandHandler := handlers.Get(r.Form.Get("command"))
	result := commandHandler.Process(r.Form.Get("text"), slackUser)

	response := map[string]string{
		"text":          result.Content,
		"response_type": configs.Commands.ResponseType,
	}

	jsonResponse, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
