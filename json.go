package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// responseWithJSON take the input as a Go data structure type and output a json
func responseWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	data, err := json.Marshal(payload)

	// Handle error if Marshal process gone wrong
	if err != nil {
		log.Printf("Failed to Marshal json response: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, writeErr := w.Write(data)
	if writeErr != nil {
		log.Printf("Failed while writing data after Marshal: %v", err)
		w.WriteHeader(500)
		return
	}
}

func responseWithError(w http.ResponseWriter, statusCode int, msg string) {
	// Handle when back end error
	if statusCode > 499 {
		log.Println("Responding with 5XX error", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	responseWithJSON(w, statusCode, errResponse{
		Error: msg,
	})
}
