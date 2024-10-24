package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, status int, message string) {

	if status > 499 {
		log.Printf("Error code 5XX, server error on our end %v", message)
	}

	type errorMessage struct {
		Error string `json:"Error"`
	}

	respondWithJSON(w, status, errorMessage{
		Error: message,
	})
}

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
	return
}
