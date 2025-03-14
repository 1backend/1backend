package sdk

import openapi "github.com/openorch/openorch/clients/go"

type ClientFactory interface {
	Client(opts ...ClientOption) *openapi.APIClient
}
