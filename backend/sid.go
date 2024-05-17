package main

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func GenerateIDAndSID(r *http.Request, siteID string) (string, string) {
	h := sha1.New()

	h.Write([]byte(Config.SessionSalt)) // Unique globally
	h.Write([]byte(siteID))             // Unique per site
	h.Write([]byte(GetIPAddress(r)))    // Unique per IP
	h.Write([]byte(r.UserAgent()))      // Unique per user-agent
	sid := fmt.Sprintf("%x", h.Sum(nil))

	h.Write([]byte(r.URL.Path))
	h.Write([]byte(time.Now().Format(time.RFC3339Nano)))
	id := fmt.Sprintf("%x", h.Sum(nil))

	return id, sid
}

func GetIPAddress(r *http.Request) string {
	ipAndPort := r.RemoteAddr
	if forwardedFor := r.Header.Get("X-Forwarded-For"); forwardedFor != "" {
		// Leftmost IP in IP,IP,IP is the originating client
		ipAndPort = strings.Split(forwardedFor, ",")[0]
	}

	onlyIP := strings.Split(ipAndPort, ":")[0]
	return onlyIP
}
