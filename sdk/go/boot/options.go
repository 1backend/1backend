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
	Test      bool
	ServerUrl string
	SelfUrl   string

	TokenAutoRefreshOff bool
	ClientFactory       client.ClientFactory
	TokenRefresher      endpoint.TokenRefresher
	Middlewares         func(http.HandlerFunc) http.HandlerFunc
	Authorizer          auth.Authorizer
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
		o.ClientFactory = client.NewApiClientFactory(o.SelfUrl)
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
