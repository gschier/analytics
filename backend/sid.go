package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/sebest/xff"
	"net"
	"net/http"
	"time"
)

func GenerateIDAndSID(r *http.Request, siteID string) (string, string) {
	return GenerateIDAndSIDForIP(r, siteID, GetIdentityIPAddress(r, siteID))
}

func GenerateIDAndSIDForIP(r *http.Request, siteID, ipAddress string) (string, string) {
	h := sha1.New()

	h.Write([]byte(Config.SessionSalt)) // Unique globally
	h.Write([]byte(siteID))             // Unique per site
	h.Write([]byte(ipAddress))          // Unique per IP
	h.Write([]byte(r.UserAgent()))      // Unique per user-agent
	sid := fmt.Sprintf("%x", h.Sum(nil))

	h.Write([]byte(r.URL.Path))
	h.Write([]byte(time.Now().Format(time.RFC3339Nano)))
	id := fmt.Sprintf("%x", h.Sum(nil))

	return id, sid
}

func GetIdentityIPAddress(r *http.Request, siteID string) string {
	if ip := GetTrustedForwardedIPAddress(r, siteID); ip != "" {
		return ip
	}
	return GetIPAddress(r)
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

func GetTrustedForwardedIPAddress(r *http.Request, siteID string) string {
	if Config.AnalyticsForwardSecret == "" {
		return ""
	}

	ipAddress := r.Header.Get("X-Analytics-Origin")
	if ipAddress == "" {
		ipAddress = r.Header.Get("X-Analytics-Client-IP")
	}
	signature := r.Header.Get("X-Analytics-Signature")
	if ipAddress == "" || signature == "" || net.ParseIP(ipAddress) == nil {
		return ""
	}

	expected := SignAnalyticsForwardedIP(Config.AnalyticsForwardSecret, siteID, ipAddress, r.UserAgent())
	if !hmac.Equal([]byte(signature), []byte(expected)) {
		NewLogger("ip").WarnContext(r.Context(), "Invalid analytics forwarded IP signature")
		return ""
	}

	return ipAddress
}

func SignAnalyticsForwardedIP(secret, siteID, ipAddress, userAgent string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(siteID))
	mac.Write([]byte{0})
	mac.Write([]byte(ipAddress))
	mac.Write([]byte{0})
	mac.Write([]byte(userAgent))
	return hex.EncodeToString(mac.Sum(nil))
}
