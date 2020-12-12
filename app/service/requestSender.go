package service

import (
	"net/http"
	"time"

	"github.com/onurkybsi/rester/app/model"
)

var tr = &http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    30 * time.Second,
	DisableCompression: true,
}

var client = &http.Client{Transport: tr}

// Ping : Send ping for checking server status.
func Ping(targetServerURL string) (int, error) {
	req, err := http.NewRequest(model.HeadMethod, targetServerURL, nil)

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
