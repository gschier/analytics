package main

import (
	"encoding/json"
	"net/http"
)

func IpOrTzToCountryCode(ip, tz string) string {
	countryFromTz := TimezoneToCountryCode(tz)
	resp, err := http.Get("http://ip-api.com/json/" + ip)
	if err != nil {
		NewLogger("ip").Warn("Failed to fetch ip info. Falling back to TZ", "error", err)
		return countryFromTz
	}
	defer resp.Body.Close()

	var ipResp struct {
		CountryCode string `json:"countryCode"`
	}
	err = json.NewDecoder(resp.Body).Decode(&ipResp)
	if err != nil {
		NewLogger("ip").Warn("Failed to fetch ip info. Falling back to TZ", "error", err)
		return countryFromTz
	}

	return ipResp.CountryCode
}
