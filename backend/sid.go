package main

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"strings"
)

func GenerateSID(r *http.Request, siteID string) string {
	sid := sha1.New()

	sid.Write([]byte(Config.SessionSalt)) // Unique globally
	sid.Write([]byte(siteID))             // Unique per site
	sid.Write([]byte(GetIPAddress(r)))    // Unique per IP
	sid.Write([]byte(r.UserAgent()))      // Unique per user-agent

	return fmt.Sprintf("%x", sid.Sum(nil))
}

func GetIPAddress(r *http.Request) string {
	ipAndPort := r.RemoteAddr
	if forwardedFor := r.Header.Get("X-Forwarded-For"); forwardedFor != "" {
		ipAndPort = forwardedFor
	}

	return strings.Split(ipAndPort, ":")[0]
}
