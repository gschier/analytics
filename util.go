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
		_, _ = w.Write([]byte("Failed to encode output"))
	}
}
