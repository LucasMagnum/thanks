package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)


func main() {
    fmt.Println("Starting listening 0.0.0.0:4390")
    http.HandleFunc("/", handleSlackCommand)
    http.ListenAndServe(":4390", nil)
}


func handleSlackCommand(w http.ResponseWriter, r *http.Request) {
    response := map[string]string{
        "text": "Hello world feedback <3",
        "response_type": "in_channel",
    }

    jsonResponse, _ := json.Marshal(response)

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)
}