package main

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func GenerateSID(r *http.Request, siteID string) string {
	ipAndPort := r.RemoteAddr
	if forwardedFor := r.Header.Get("X-Forwarded-For"); forwardedFor != "" {
		ipAndPort = forwardedFor
	}
	ip := strings.Split(ipAndPort, ":")[0]

	n := time.Now()
	dayStart := time.Date(n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, n.Location())

	sid := sha256.New()
	sid.Write([]byte(Config.SessionSalt))            // Unique globally
	sid.Write([]byte(siteID))                        // Unique per site
	sid.Write([]byte(dayStart.Format(time.RFC3339))) // Unique per day
	sid.Write([]byte(ip))                            // Unique per IP
	sid.Write([]byte(r.UserAgent()))                 // Unique per user-agent

	return fmt.Sprintf("%x", sid.Sum(nil))
}
