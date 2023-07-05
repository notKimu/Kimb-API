package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Panicln("[ERROR] Server responded with 5xx error, valve pls fix >>>\n", message)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	/*
		{
			"error": "message"
		}
	*/
	respondWithJSON(w, code, errorResponse{
		Error: message,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("[ERROR] Failed to marshal JSON response >>>/n%v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
