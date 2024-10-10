package http

import (
	"crypto/tls"
	"net/http"
)

func NewHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,

			}
		}
	}
}
