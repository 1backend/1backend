/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package boot

import (
	"net/http"
	"os"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/middlewares"
	"github.com/pkg/errors"
)

type Options struct {
	Test bool

	// ServerUrl is the URL of the 1Backend server.
	ServerUrl string

	// SelfUrl is the URL of the service itself.
	// Used for service registration.
	SelfUrl string

	// If set to true, expired tokens won't be autorefreshed by
	// the server.
	TokenAutoRefreshOff bool

	// ClientFactory is used for service to service communication
	// ie. this is how services call each other
	ClientFactory client.ClientFactory

	TokenRefresher    endpoint.TokenRefresher
	PermissionChecker endpoint.PermissionChecker
	Middlewares       func(http.HandlerFunc) http.HandlerFunc

	// Authorizer is a helper interface that contains
	// auth related utility functions
	Authorizer auth.Authorizer
}

func (o *Options) LoadEnvars() error {
	if o.ServerUrl == "" {
		o.ServerUrl = os.Getenv("OB_SERVER_URL")
	}

	if o.ServerUrl == "" {
		o.ServerUrl = "http://127.0.0.1:11337"
	}

	if o.SelfUrl == "" {
		o.SelfUrl = os.Getenv("OB_SELF_URL")
	}

	if !o.Test && os.Getenv("OB_TEST") == "true" {
		o.Test = true
	}

	if o.ClientFactory == nil {
		o.ClientFactory = client.NewApiClientFactory(o.ServerUrl)
	}

	if o.Authorizer == nil {
		o.Authorizer = auth.AuthorizerImpl{}
	}

	if o.TokenRefresher == nil {
		var err error
		o.TokenRefresher, err = endpoint.NewTokenRefresher(
			o.ClientFactory,
			o.Authorizer,
		)
		if err != nil {
			return errors.Wrap(err, "failed to create token refresher")
		}
	}

	if o.PermissionChecker == nil {
		o.PermissionChecker = endpoint.NewPermissionChecker(
			o.ClientFactory,
		)
	}

	if os.Getenv("OB_TOKEN_AUTO_REFRESH_OFF") == "true" {
		o.TokenAutoRefreshOff = true
	}

	if o.Middlewares == nil {
		mws := middlewares.Applicator(
			o.TokenRefresher,
			o.TokenAutoRefreshOff,
		)
		o.Middlewares = mws
	}

	return nil
}
