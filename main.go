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
		type bucket struct {
			Count   int
			Percent int
		}
		events, _ := db.ListAnalyticsEvents(r.Context())
		buckets := make([]bucket, 50)

		bucketDuration := time.Minute

		for n := 0; n < len(buckets); n++ {
			nd := time.Duration(n)
			now := time.Now()
			nowRounded := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, now.Location())
			start := nowRounded.Add(-(nd + 1) * bucketDuration)
			end := nowRounded.Add(-nd * bucketDuration)

			for _, e := range events {
				if e.CreatedAt.After(start) && e.CreatedAt.Before(end) {
					buckets[n].Count += 1
				}
			}
		}

		maxCount := 0
		for _, b := range buckets {
			if b.Count > maxCount {
				maxCount = b.Count
			}
		}

		for i := range buckets {
			buckets[i].Percent = int(float64(buckets[i].Count) / float64(maxCount) * 100)
		}

		reverseBuckets := make([]bucket, len(buckets))
		for i := range buckets {
			reverseBuckets[i] = buckets[len(buckets)-i-1]
		}

		RenderTemplate(w, "index.gohtml", map[string]interface{}{
			"Title":   "Analytics",
			"Events":  events,
			"Buckets": reverseBuckets,
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
