package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func SetupRouter(globalWebsiteID string) http.Handler {
	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	r.Path("/script.js").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/tracker.js")
	})

	r.Path("/").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pageviews := FindAnalyticsPageviews(GetDB(), r.Context(), globalWebsiteID)
		RenderTemplate(r, w, "index.gohtml", "Analytics", map[string]interface{}{
			"Events":  pageviews,
			"Buckets": RollupPageviews(time.Now().Add(-23*time.Hour), 24, PeriodHour, pageviews),
		})
	})

	r.Path("/events").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		events := FindAnalyticsEvents(GetDB(), r.Context(), globalWebsiteID)
		RespondJSON(w, events)
	})

	r.Path("/event").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

		site := q.Get("id")
		eventName := q.Get("e")
		sid := GenerateSID(r, site)

		CreateAnalyticsEvent(GetDB(), r.Context(), site, eventName, sid)
	})

	r.Path("/page").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

		site := q.Get("id")
		path := q.Get("p")
		host := q.Get("h")
		screensize := q.Get("xy")
		timezone := q.Get("tz")

		userAgent := r.UserAgent()
		sid := GenerateSID(r, site)
		countryCode := TimezoneToCountryCode[timezone]

		CreateAnalyticsPageview(GetDB(), r.Context(), site, host, path, screensize, countryCode, sid, userAgent)
	})

	return r
}
