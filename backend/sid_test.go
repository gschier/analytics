package main

import (
	"net/http"
	"testing"
)

func TestGetAnalyticsOriginIPAddress(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/t/e?id=site_123", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-Analytics-Origin", "203.0.113.5")

	if got := GetAnalyticsOriginIPAddress(req); got != "203.0.113.5" {
		t.Fatalf("expected analytics origin IP, got %q", got)
	}
}

func TestGetAnalyticsOriginIPAddressRejectsInvalidIP(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/t/e?id=site_123", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-Analytics-Origin", "invalid")

	if got := GetAnalyticsOriginIPAddress(req); got != "" {
		t.Fatalf("expected invalid analytics origin IP to be rejected, got %q", got)
	}
}
