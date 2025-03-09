package main

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"time"
)

func SetupRouter() http.Handler {
	r := mux.NewRouter()

	r.Path("/robots.txt").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("User-agent: *\nDisallow: /"))
	})

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

	r.Path("/api/popular_events").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siteId := r.URL.Query().Get("site")
		counts := FindAnalyticsEventsPopular(
			GetDB(),
			r.Context(),
			start(),
			end(),
			siteId,
		)
		RespondJSON(w, &counts)
	})

	r.Path("/api/popular_referrers").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siteId := r.URL.Query().Get("site")
		counts := FindAnalyticsReferrersPopular(
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

	r.Path("/t/e").Methods(http.MethodGet, http.MethodPost).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()

		site := r.Form.Get("id")
		eventName := r.Form.Get("e")
		attributes := r.Form.Get("a")
		screensize := r.Form.Get("xy")
		timezone := r.Form.Get("tz")
		platform := r.Form.Get("os")
		version := r.Form.Get("v")
		id, sid := GenerateIDAndSID(r, site)

		uid := r.Form.Get("u")
		if uid == "" {
			uid = sid
		}

		if attributes == "" {
			attributes = "{}"
		}

		event := AnalyticsEvent{
			ID:          id,
			SID:         sid,
			UID:         uid,
			WebsiteID:   site,
			Name:        eventName,
			Attributes:  attributes,
			CountryCode: IpOrTzToCountryCode(GetIPAddress(r), timezone),
			ScreenSize:  screensize,
			Platform:    platform,
			Version:     version,
		}
		CreateAnalyticsEvent(GetDB(), r.Context(), &event)
		RespondText(w, "OK")
	})

	r.Path("/t/p").Methods(http.MethodGet, http.MethodPost).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()

		path := r.Form.Get("p")
		site := r.Form.Get("id")
		host := r.Form.Get("h")
		screensize := r.Form.Get("xy")
		timezone := r.Form.Get("tz")
		referrer := r.Form.Get("r")
		id, sid := GenerateIDAndSID(r, site)

		uid := r.Form.Get("u")
		if uid == "" {
			uid = sid
		}

		// Sanitize path
		path = strings.TrimSuffix(path, "/")
		if path == "" {
			path = "/"
		}

		// Sanitize referrer
		referrer = strings.TrimSuffix(referrer, "/")

		pageview := AnalyticsPageview{
			ID:          id,
			SID:         sid,
			UID:         uid,
			WebsiteID:   site,
			Host:        host,
			Path:        path,
			ScreenSize:  screensize,
			CountryCode: TimezoneToCountryCode(timezone),
			UserAgent:   r.UserAgent(),
			Referrer:    referrer,
		}
		CreateAnalyticsPageview(GetDB(), r.Context(), &pageview)
		RespondText(w, "OK")
	})

	r.PathPrefix("/assets").Handler(http.FileServer(http.Dir("./dist")))

	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./dist/index.html")
	})

	return r
}

func start() time.Time {
	return time.Now().Add(-30 * PeriodDay)
}

func end() time.Time {
	return time.Now()
}
