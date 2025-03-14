package sdk

import (
	"net/http"
	"strings"

	openapi "github.com/openorch/openorch/clients/go"
)

type ClientOption func(*openapi.Configuration)

func WithToken(token string) ClientOption {
	return func(cfg *openapi.Configuration) {
		cfg.DefaultHeader = map[string]string{
			"Authorization": "Bearer " + token,
		}
	}
}

func WithTokenFromRequest(req *http.Request) ClientOption {
	authHeader := req.Header.Get("Authorization")
	authHeader = strings.Replace(authHeader, "Bearer ", "", 1)

	return func(cfg *openapi.Configuration) {
		cfg.DefaultHeader = map[string]string{
			"Authorization": "Bearer " + authHeader,
		}
	}
}

func WithAddress(address string) ClientOption {
	return func(cfg *openapi.Configuration) {
		cfg.Servers = openapi.ServerConfigurations{
			{
				URL:         address,
				Description: "Default server",
			},
		}
	}
}

func CustomHeader(key, value string) ClientOption {
	return func(cfg *openapi.Configuration) {
		if cfg.DefaultHeader == nil {
			cfg.DefaultHeader = make(map[string]string)
		}
		cfg.DefaultHeader[key] = value
	}
}

type APIClientFactory struct {
	url string
}

func NewApiClientFactory(url string) *APIClientFactory {
	return &APIClientFactory{
		url: url,
	}
}

func (f *APIClientFactory) Client(opts ...ClientOption) *openapi.APIClient {
	cfg := &openapi.Configuration{
		Servers: openapi.ServerConfigurations{
			{
				URL:         f.url,
				Description: "Default server",
			},
		},
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return openapi.NewAPIClient(cfg)
}
