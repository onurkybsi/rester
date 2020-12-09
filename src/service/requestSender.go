package service

import (
	"net/http"
	"time"
)

var tr = &http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    30 * time.Second,
	DisableCompression: true,
}

var client = &http.Client{Transport: tr}

// Ping : Send ping for checking server status.
func Ping(domain string) (int, error) {
	url := "http://" + domain

	req, err := http.NewRequest("HEAD", url, nil)

	if err != nil {
		return 0, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return 0, err

	}

	resp.Body.Close()

	return resp.StatusCode, nil
}
