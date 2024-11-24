package main

import (
	"crypto/sha1"
	"fmt"
	"github.com/sebest/xff"
	"net"
	"net/http"
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
	addr := xff.GetRemoteAddr(r)
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		NewLogger("ip").WarnContext(r.Context(), "Failed to split host and port", "error", err)
		return addr
	}

	return host
}
