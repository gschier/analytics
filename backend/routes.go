package main

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func SetupRouter() http.Handler {
	r := mux.NewRouter()

	r.Path("/script.js").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=86400")
		http.ServeFile(w, r, "./dist/tracker.js")
	})

	r.Path("/api/websites").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		account, _ := GetAccountByEmail(GetDB(), context.Background(), "greg@schier.co")
		websites := FindWebsitesByAccountID(
			GetDB(),
			r.Context(),
			account.ID,
		)
		RespondJSON(w, &websites)
	})

	r.Path("/api/rollups/pageviews").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siteId := r.URL.Query().Get("site")
		rollups := FindAnalyticsPageviewsBuckets(
			GetDB(),
			r.Context(),
			start(),
			end(),
			PeriodDay,
			siteId,
		)
		RespondJSON(w, &rollups)
	})

	r.Path("/api/popular").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siteId := r.URL.Query().Get("site")
		counts := FindAnalyticsPageviewsPopular(
			GetDB(),
			r.Context(),
			start(),
			end(),
			siteId,
		)
		RespondJSON(w, &counts)
	})

	r.Path("/api/live").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siteId := r.URL.Query().Get("site")
		count := CountAnalyticsPageviewsRecent(GetDB(), r.Context(), siteId)
		RespondJSON(w, &count)
	})

	r.Path("/t/e").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

		site := q.Get("id")
		eventName := q.Get("e")
		attributes := q.Get("a")
		id, sid := GenerateIDAndSID(r, site)

		CreateAnalyticsEvent(GetDB(), r.Context(), id, sid, site, eventName, attributes)
	})

	r.Path("/t/p").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

		site := q.Get("id")
		path := q.Get("p")
		host := q.Get("h")
		screensize := q.Get("xy")
		timezone := q.Get("tz")

		userAgent := r.UserAgent()
		id, sid := GenerateIDAndSID(r, site)
		countryCode := TimezoneToCountryCode(timezone)
		CreateAnalyticsPageview(GetDB(), r.Context(), id, sid, site, host, path, screensize, countryCode, userAgent)
		RespondText(w, "OK")
	})

	r.PathPrefix("/assets").Handler(http.FileServer(http.Dir("./dist")))

	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./dist/index.html")
	})

	return r
}

func start() time.Time {
	return time.Now().Add(-60 * PeriodDay)
}

func end() time.Time {
	return time.Now()
}
