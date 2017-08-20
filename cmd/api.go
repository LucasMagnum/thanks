package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lucasmagnum/thanks/pkg/api"
)

func main() {
	port := fmt.Sprintf(":%d", 8080)

	log.Printf("Starting listening 0.0.0.0%s", port)

	http.HandleFunc("/feedback", api.FeedbackHandler)
	//http.HandleFunc("/raking", api.RankingHandler)

	http.ListenAndServe(port, nil)
}
