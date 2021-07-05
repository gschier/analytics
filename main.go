package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println()
	fmt.Printf("\u001B[32;1m┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓\u001B[0m\n")
	fmt.Printf("\u001B[32;1m┃                  analytics                  ┃\u001B[0m\n")
	fmt.Printf("\u001B[32;1m┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛\u001B[0m\n")

	db := NewDB()
	if Config.MigrateOnStart {
		mustMigrate(context.Background(), db.db)
	}

	handler := mux.NewRouter()
	handler.Path("/").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		events, _ := db.ListEvents(r.Context())
		RespondJSON(w, events)
	})

	handler.Path("/page").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		events, _ := db.ListEvents(r.Context())
		RespondJSON(w, events)
	})

	fmt.Println("[schier.co] \033[32;1mStarted server on http://" + Config.Host + ":" + Config.Port + "\033[0m")
	log.Fatal(http.ListenAndServe(Config.Host+":"+Config.Port, handler))
}
