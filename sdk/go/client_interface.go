package sdk

import openapi "github.com/1backend/1backend/clients/go"

type ClientFactory interface {
	Client(opts ...ClientOption) *openapi.APIClient
}
