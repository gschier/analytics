package main

import (
	"encoding/json"
	"net/http"
)

var _ipCache = make(map[string]string)

func IpOrTzToCountryCode(ip, tz string) string {
	countryFromTz := TimezoneToCountryCode(tz)
	code, err := IpToCountryCode(ip)
	if err != nil {
		NewLogger("ip").Warn("Failed to fetch ip info. Falling back to TZ", "error", err)
		return countryFromTz
	}
	return code
}

func IpToCountryCode(ip string) (string, error) {
	if v, ok := _ipCache[ip]; ok {
		return v, nil
	}

	resp, err := http.Get("http://ip-api.com/json/" + ip)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var ipResp struct {
		CountryCode string `json:"countryCode"`
	}
	err = json.NewDecoder(resp.Body).Decode(&ipResp)
	if err != nil {
		return "", err
	}

	_ipCache[ip] = ipResp.CountryCode
	return ipResp.CountryCode, nil
}
