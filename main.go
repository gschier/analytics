package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println()
	fmt.Printf("\u001B[32;1m┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓\u001B[0m\n")
	fmt.Printf("\u001B[32;1m┃                  analytics                  ┃\u001B[0m\n")
	fmt.Printf("\u001B[32;1m┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛\u001B[0m\n")

	InitConfig()

	db := NewDB()
	if Config.MigrateOnStart {
		mustMigrate(context.Background(), db.db)
	}

	globalWebsiteID := ""
	account, err := db.GetAccountByEmail(context.Background(), "greg@schier.co")
	if err == sql.ErrNoRows {
		a, _ := db.CreateAccount(context.Background(), "greg@schier.co", "my-pass!")
		w, _ := db.CreateWebsite(context.Background(), a.ID, "My Blog")
		globalWebsiteID = w.ID
		println("WEBSITE:", w.ID)
	} else if err != nil {
		panic(err)
	} else {
		websites, _ := db.FindWebsitesByAccountID(context.Background(), account.ID)
		globalWebsiteID = websites[0].ID
	}
	println("WEBSITE:", globalWebsiteID)

	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	r.Path("/").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pageviews, err := db.FindAnalyticsPageviews(r.Context(), globalWebsiteID)
		if err != nil {
			RespondError(w, err)
			return
		}

		// Shift one minute into the future to capture latest incomplete bucket
		buckets := RollupPageviews(time.Now().Add(-time.Hour+time.Minute), 60, PeriodMinute, pageviews)

		RenderTemplate(r, w, "index.gohtml", map[string]interface{}{
			"Title":     "Analytics",
			"Events":    pageviews,
			"Buckets":   buckets,
			"WebsiteID": globalWebsiteID,
		})
	})

	r.Path("/events").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		events, _ := db.FindAnalyticsEvents(r.Context(), globalWebsiteID)
		RespondJSON(w, events)
	})

	r.Path("/event").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		websiteID := r.URL.Query().Get("id")
		eventName := r.URL.Query().Get("e")

		_, err := db.CreateAnalyticsEvent(r.Context(), websiteID, eventName)
		if err != nil {
			RespondError(w, err)
			return
		}
	})

	r.Path("/page").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		site := r.URL.Query().Get("id")
		path := r.URL.Query().Get("p")
		host := r.URL.Query().Get("h")
		screensize := r.URL.Query().Get("xy")
		timezone := r.URL.Query().Get("tz")
		countryCode := TimezoneToCountryCode[timezone]
		sid := GenerateSID(r, site)

		_, err := db.CreateAnalyticsPageview(r.Context(), site, host, path, screensize, countryCode, sid)
		if err != nil {
			RespondError(w, err)
			return
		}
	})

	fmt.Println("[schier.co] \033[32;1mStarted server on http://" + Config.Host + ":" + Config.Port + "\033[0m")
	log.Fatal(http.ListenAndServe(Config.Host+":"+Config.Port, r))
}
