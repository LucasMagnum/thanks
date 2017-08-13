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

// HandleSlackCommand handle the request and execute the configured
// handler for the command received
func HandleSlackCommand(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if !isDomainAllowed(r.Form.Get("team_domain")) {
		http.Error(w, "Domain requests not allowed", 403)
		return
	}

	slackUser := handlers.SlackUser{
		UserId:   r.Form.Get("user_id"),
		Username: r.Form.Get("user_name"),
	}

	log.Println("Request received from", r.Form.Get("team_domain"))
	log.Println("Command: ", r.Form.Get("command"))
	log.Println("Text: ", r.Form.Get("text"))
	log.Println("Request UserId: ", r.Form.Get("user_id"))
	log.Println("Request Username: ", r.Form.Get("user_name"))

	commandHandler := handlers.Get(r.Form.Get("command"))
	result := commandHandler.Process(r.Form.Get("text"), slackUser)

	log.Println("Returning message:", result.Content)
	response := map[string]string{
		"text":          result.Content,
		"response_type": configs.Commands.ResponseType,
	}

	jsonResponse, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// isDomainAllowed should return true when API.CheckDomain is active
// and the teamDomain is equal to API.AllowdTeamDomain
func isDomainAllowed(teamDomain string) bool {
	if configs.API.CheckDomain && configs.API.AllowedTeamDomain != teamDomain {
		return false
	}
	return true
}
