package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func SetupRouter() http.Handler {
	r := mux.NewRouter()

	r.Path("/script.js").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./dist/tracker.js")
	})

	r.Path("/api/rollups/pageviews").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rollups := FindAnalyticsPageviewsHourly(
			GetDB(),
			r.Context(),
			time.Now().Add(-24*7*time.Hour+time.Hour),
			time.Now().Add(time.Hour),
			ensureDummyWebsite(),
		)
		RespondJSON(w, &rollups)
	})

	r.Path("/api/pageviews").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pageviews := FindAnalyticsPageviews(GetDB(), r.Context(), ensureDummyWebsite())
		RespondJSON(w, &pageviews)
	})

	r.Path("/api/event").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

		site := q.Get("id")
		eventName := q.Get("e")
		id, sid := GenerateIDAndSID(r, site)

		CreateAnalyticsEvent(GetDB(), r.Context(), id, sid, site, eventName)
	})

	r.Path("/api/page").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
