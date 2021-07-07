package main

import (
	"context"
	"database/sql"
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

	account, err := db.GetAccountByEmail(context.Background(), "greg@schier.co")
	if err == sql.ErrNoRows {
		a, _ := db.CreateAccount(context.Background(), "greg@schier.co", "my-pass!")
		w, _ := db.CreateWebsite(context.Background(), a.ID, "My Blog")
		println("WEBSITE:", w.ID)
	} else if err != nil {
		panic(err)
	} else {
		websites, _ := db.FindWebsitesByAccountID(context.Background(), account.ID)
		println("WEBSITE:", websites[0].ID)
	}

	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	r.Path("/").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RenderTemplate(w, "index.gohtml", map[string]interface{}{
			"Title": "Analytics",
		})
	})

	r.Path("/events").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		events, _ := db.ListAnalyticsEvents(r.Context())
		RespondJSON(w, events)
	})

	r.Path("/event").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		websiteID := r.URL.Query().Get("website")
		event, err := db.CreateAnalyticsEvent(r.Context(), websiteID, name)
		if err != nil {
			RespondError(w, err)
			return
		}

		RespondJSON(w, event)
	})

	fmt.Println("[schier.co] \033[32;1mStarted server on http://" + Config.Host + ":" + Config.Port + "\033[0m")
	log.Fatal(http.ListenAndServe(Config.Host+":"+Config.Port, r))
}
