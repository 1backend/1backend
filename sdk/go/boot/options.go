/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package boot

import (
	"context"
	"net/http"
	"os"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/infra"
	"github.com/1backend/1backend/sdk/go/middlewares"
	"github.com/1backend/1backend/sdk/go/telemetry"
	"github.com/pkg/errors"
)

type Option func(*Options)

type Options struct {
	Test bool

	// ServiceName is the service identity used for telemetry and other
	// service-level defaults. Use NewOptions to set this for normal services.
	ServiceName string

	// ServerUrl is the internally addressable URL of the 1Backend server.
	// It configures the API client used for service-to-service calls.
	// Defaults to OB_INTERNAL_SERVER_URL, then deprecated OB_SERVER_URL.
	ServerUrl string

	// SelfUrl is the public URL of the service itself.
	// It is used for service registration.
	// Defaults to OB_PUBLIC_URL, then deprecated OB_SELF_URL.
	SelfUrl string

	// Db is the database engine (for example: postgres, mysql).
	Db string

	// DbConnectionString is the database connection string.
	DbConnectionString string

	// ReadDbConnectionString is the optional read-replica database connection string.
	ReadDbConnectionString string

	// DbApplicationName is used as the base PostgreSQL application_name label.
	DbApplicationName string

	// DbPool controls database/sql pool sizing and connection lifetime.
	DbPool infra.DbPoolConfig

	// If set to true, expired tokens won't be autorefreshed by
	// the server.
	TokenAutoRefreshOff bool

	// ClientFactory is used for service to service communication
	// ie. this is how services call each other
	ClientFactory client.ClientFactory

	TokenRefresher    endpoint.TokenRefresher
	PermissionChecker endpoint.PermissionChecker
	TokenExchanger    endpoint.TokenExchanger
	Middlewares       func(http.HandlerFunc) http.HandlerFunc

	// Authorizer is a helper interface that contains
	// auth related utility functions
	Authorizer auth.Authorizer

	// Telemetry configures the default OpenTelemetry setup. It is enabled by
	// default when ServiceName is set, and can be disabled with
	// WithTelemetryDisabled or OB_OTEL_DISABLED=true.
	Telemetry telemetry.Config

	// TelemetrySetup overrides the telemetry setup function. This is mainly
	// useful for tests and custom hosting environments.
	TelemetrySetup func(context.Context, telemetry.Config) (telemetry.ShutdownFunc, string, error)

	telemetryShutdown    telemetry.ShutdownFunc
	telemetryMetricsPath string
	telemetryInitialized bool
}

// NewOptions returns the default boot options for a service. The default set
// includes OpenTelemetry setup; LoadEnvars starts it automatically before
// datastore factories are created.
func NewOptions(serviceName string, opts ...Option) *Options {
	o := &Options{
		ServiceName: serviceName,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(o)
		}
	}
	return o
}

func WithServerUrl(serverURL string) Option {
	return func(o *Options) {
		o.ServerUrl = serverURL
	}
}

func WithSelfUrl(selfURL string) Option {
	return func(o *Options) {
		o.SelfUrl = selfURL
	}
}

func WithTelemetryConfig(cfg telemetry.Config) Option {
	return func(o *Options) {
		o.Telemetry = cfg
	}
}

func WithTelemetryDisabled() Option {
	return func(o *Options) {
		o.Telemetry.Disabled = true
	}
}

func WithTelemetrySetup(setup func(context.Context, telemetry.Config) (telemetry.ShutdownFunc, string, error)) Option {
	return func(o *Options) {
		o.TelemetrySetup = setup
	}
}

func (o *Options) LoadEnvars() error {
	if err := o.loadEnvars(); err != nil {
		return err
	}

	_, _, err := o.ensureTelemetry(context.Background(), false)
	return err
}

func (o *Options) loadEnvars() error {
	if o.ServerUrl == "" {
		o.ServerUrl = os.Getenv("OB_INTERNAL_SERVER_URL")
	}

	if o.ServerUrl == "" {
		o.ServerUrl = os.Getenv("OB_SERVER_URL")
	}

	if o.ServerUrl == "" {
		o.ServerUrl = "http://127.0.0.1:11337"
	}

	if o.SelfUrl == "" {
		o.SelfUrl = os.Getenv("OB_PUBLIC_URL")
	}

	if o.SelfUrl == "" {
		o.SelfUrl = os.Getenv("OB_SELF_URL")
	}

	if !o.Test && os.Getenv("OB_TEST") == "true" {
		o.Test = true
	}

	if o.Db == "" {
		o.Db = os.Getenv("OB_DB")
	}

	if o.DbConnectionString == "" {
		o.DbConnectionString = os.Getenv("OB_DB_CONNECTION_STRING")
	}

	if o.ReadDbConnectionString == "" {
		o.ReadDbConnectionString = os.Getenv("OB_DB_READ_CONNECTION_STRING")
	}

	if o.DbApplicationName == "" {
		o.DbApplicationName = os.Getenv("OB_DB_APPLICATION_NAME")
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

	if o.TokenExchanger == nil {
		o.TokenExchanger = endpoint.NewTokenExchanger(
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
