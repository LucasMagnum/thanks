package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/lucasmagnum/thanks/pkg"
)

func main() {
	port := fmt.Sprintf(":%d", 8080)

	log.Printf("Starting listening 0.0.0.0%s", port)

	http.HandleFunc("/feedback", handleFeedbackCommand)
	http.HandleFunc("/ranking", handleRankingCommand)

	http.ListenAndServe(port, nil)
}

func handleFeedbackCommand(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	handler := app.NewFeedbackHandler()
	command := app.NewCommand(
		r.Form.Get("text"),
		r.Form.Get("user_id"),
		r.Form.Get("user_name"),
	)

	responseText, err := handler.ProcessCommand(command)

	responseType := "in_channel"
	if err != nil {
		responseText = err.Error()
		responseType = "ephemeral"
	}

	jsonResponse, _ := json.Marshal(map[string]string{
		"text":          responseText,
		"response_type": responseType,
	})

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func handleRankingCommand(w http.ResponseWriter, r *http.Request) {
	jsonResponse, _ := json.Marshal(map[string]string{
		"text":          "A.W.E.S.O.M.E",
		"response_type": "in_channel",
	})

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
