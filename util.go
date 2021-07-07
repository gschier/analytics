package main

import (
	"encoding/json"
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

func RespondError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte(err.Error()))
}
