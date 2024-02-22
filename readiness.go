package main

import (
	"log"
	"net/http"
)

func (cfg *ApiConfig) handleReadiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	log.Fatal(w.Write([]byte(http.StatusText(http.StatusOK))))
}
