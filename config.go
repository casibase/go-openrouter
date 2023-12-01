package openrouter

import (
	"net/http"
)

const (
	routerAPIURLv1                 = "https://openrouter.ai/api/v1"
	defaultEmptyMessagesLimit uint = 300
)

// ClientConfig is a configuration of a client.
// XTitle、HttpRefer your own site url
type ClientConfig struct {
	authToken          string
	XTitle             string
	HttpReferer        string
	BaseURL            string
	HTTPClient         *http.Client
	EmptyMessagesLimit uint
}

func DefaultConfig(auth, xTitle, httpReferer string) (ClientConfig, error) {
	return ClientConfig{
		authToken:          auth,
		HTTPClient:         &http.Client{},
		XTitle:             xTitle,
		HttpReferer:        httpReferer,
		BaseURL:            routerAPIURLv1,
		EmptyMessagesLimit: defaultEmptyMessagesLimit,
	}, nil
}

func (c ClientConfig) WithHttpClientConfig(client *http.Client) ClientConfig {
	c.HTTPClient = client
	return c
}
