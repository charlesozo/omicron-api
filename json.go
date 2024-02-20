package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, errmessage string) {
	if code < 499 {
		log.Printf("Responding with 5XX error: %s", errmessage)
	}
	type Errormsg struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, Errormsg{
		Error: errmessage,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(dat)
}
