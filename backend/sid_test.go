package main

import (
	"net/http"
	"testing"
)

func TestGetTrustedForwardedIPAddress(t *testing.T) {
	previous := Config
	t.Cleanup(func() {
		Config = previous
	})

	Config.AnalyticsForwardSecret = "secret"

	req, err := http.NewRequest(http.MethodGet, "/t/e?id=site_123", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("User-Agent", "Yaak/2026.4.0")
	req.Header.Set("X-Analytics-Client-IP", "203.0.113.5")
	req.Header.Set(
		"X-Analytics-Signature",
		SignAnalyticsForwardedIP(Config.AnalyticsForwardSecret, "site_123", "203.0.113.5", "Yaak/2026.4.0"),
	)

	if got := GetTrustedForwardedIPAddress(req, "site_123"); got != "203.0.113.5" {
		t.Fatalf("expected trusted forwarded IP, got %q", got)
	}
}

func TestGetTrustedForwardedIPAddressRejectsInvalidSignature(t *testing.T) {
	previous := Config
	t.Cleanup(func() {
		Config = previous
	})

	Config.AnalyticsForwardSecret = "secret"

	req, err := http.NewRequest(http.MethodGet, "/t/e?id=site_123", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("User-Agent", "Yaak/2026.4.0")
	req.Header.Set("X-Analytics-Client-IP", "203.0.113.5")
	req.Header.Set("X-Analytics-Signature", "invalid")

	if got := GetTrustedForwardedIPAddress(req, "site_123"); got != "" {
		t.Fatalf("expected invalid signature to be rejected, got %q", got)
	}
}
