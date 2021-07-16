package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func SetupRouter() http.Handler {
	r := mux.NewRouter()

	r.Path("/script.js").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./dist/tracker.js")
	})

	r.Path("/api/event").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

		site := q.Get("id")
		eventName := q.Get("e")
		sid := GenerateSID(r, site)

		CreateAnalyticsEvent(GetDB(), r.Context(), site, eventName, sid)
	})

	r.Path("/api/page").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./dist")))

	return r
}
