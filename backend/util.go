package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	err := enc.Encode(v)
	if err != nil {
		RespondError(w, err)
	}
}

func RespondText(w http.ResponseWriter, t string) {
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte(t))
}

func RespondError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "text/plain")
	log.Println("Internal error", err.Error())
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}
