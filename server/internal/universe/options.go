package universe

import (
	"context"
	"net/http"
	"time"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/infra"
	"github.com/1backend/1backend/sdk/go/lock"
	"github.com/1backend/1backend/sdk/go/pubsub"
	"github.com/1backend/1backend/server/internal/clients/llamacpp"
)

type Options struct {
	Port        int
	GpuPlatform string

	Az         string
	Region     string
	LLMHost    string
	VolumeName string

	// Specifies the system-wide contact email address for operational and administrative use.
	// Eg. ACME TLS certificate registration etc.
	ContactEmail string

	// Path of the config folder, configurable via the "OB_FOLDER" environment variable.
	// If Test is true, this value is ignored and a random temporary folder is used instead.
	ConfigPath string

	// BootstrapPath points at a file or directory of startup manifests to apply
	// after built-in services have started. It is configured via
	// OB_BOOTSTRAP_FOLDER.
	BootstrapPath string

	// eg. mysql, postgres
	Db string

	// Connection string eg.
	// "postgres://postgres:mysecretpassword@localhost:5432/mydatabase?sslmode=disable"
	DbConnectionString string

	// Optional read-replica connection string.
	ReadDbConnectionString string

	// Base PostgreSQL application_name label.
	DbApplicationName string

	// Database/sql pool sizing and connection lifetime.
	DbPool infra.DbPoolConfig

	// AutoIndexes enables query-observed automatic index creation for supported
	// datastore backends. It is off by default.
	AutoIndexes bool

	// Crucial for distributed features.
	// Please see the documentation for the envar OB_NODE_ID
	NodeId string

	// DbPrefix allows us to have isolated envs for different test cases
	// but still make multiple nodes in those test cases use the same
	// shard of the db.
	DbPrefix string

	SourceControlToken  string
	SecretEncryptionKey string

	FileGcs   bool
	GcpSaKey  string
	GcsBucket string

	// Url is the public, externally reachable URL of this 1Backend server.
	// It is used as the server's self identity, including OAuth/OIDC callback
	// URL construction. Defaults to OB_PUBLIC_URL, then deprecated OB_SELF_URL.
	Url string

	// ServerUrl is the internally addressable URL used by the server's own
	// API client for service-to-service calls. Defaults to OB_INTERNAL_SERVER_URL,
	// then deprecated OB_SERVER_URL, then Url.
	ServerUrl string

	// EdgeCacheItemMaxSize defines the maximum size of a single response allowed in the cache.
	// Use bytes (e.g., 50 * 1024 * 1024 for 50MB).
	EdgeCacheItemMaxSize int64

	// EdgeCacheMaxSize defines the total disk/memory footprint allowed for the file cache.
	EdgeCacheMaxSize int64

	// OB_EDGE_PROXY is used to enable the edge proxy.
	EdgeProxy bool

	EdgeProxyTestMode bool

	// Only used in tests
	EdgeProxyHttpPort int
	// Only used in tests
	EdgeProxyHttpsPort int

	// See `OB_SYNC_CERTS_TO_FILES` environment variable documentation.
	SyncCertsToFiles bool

	// See `OB_REQUIRE_VERIFIED_CONTACT` environment variable documentation.
	VerifyContacts bool

	// Test hook for provider-token verification.
	ContactAuthVerifier ContactAuthVerifier

	// Test mode if true will cause the localstore to
	// save data into random temporary folders.
	Test bool

	// Lock is a distributed lock. Use this when you want to synronize
	// across service instances/nodes.
	// eg: leader election
	Lock lock.DistributedLock

	LLamaCppClient llamacpp.ClientI

	// DataStoreFactory can create database tables
	DataStoreFactory infra.DataStoreFactory

	// PubSubFactory can create local or postgres pubsub instances.
	PubSubFactory infra.PubSubFactory

	// PubSub is the default server-level publish/subscribe instance.
	PubSub pubsub.PubSub

	// HomeDir is the 1Backend config/data/uploads/downloads directory.
	// For tests it's something like /tmp/1backend-2698538720/
	// For live it's /home/youruser/.1backend
	HomeDir string

	// ClientFactory is used for service to service communication
	// ie. this is how services call each other
	ClientFactory client.ClientFactory

	// Authorizer is a helper interface that contains
	// auth related utility functions
	Authorizer auth.Authorizer

	TokenExpiration time.Duration

	// If set to true, expired tokens won't be autorefreshed by
	// the server.
	TokenAutoRefreshOff bool

	PermissionChecker endpoint.PermissionChecker
	TokenExchanger    endpoint.TokenExchanger
	TokenRefresher    endpoint.TokenRefresher
	Middlewares       func(http.HandlerFunc) http.HandlerFunc
}

type ContactAuthProviderConfig struct {
	Id           string   `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	Kind         string   `json:"kind,omitempty"`
	IssuerURL    string   `json:"issuerUrl,omitempty"`
	ClientID     string   `json:"clientId,omitempty"`
	ClientSecret string   `json:"clientSecret,omitempty"`
	Scopes       []string `json:"scopes,omitempty"`
	GraphVersion string   `json:"graphVersion,omitempty"`
	AuthURL      string   `json:"authUrl,omitempty"`
	TokenURL     string   `json:"tokenUrl,omitempty"`
	APIURL       string   `json:"apiUrl,omitempty"`
}

type ContactAuthClaims struct {
	Provider      string
	Email         string
	EmailVerified bool
	Name          string
	Nonce         string
}

type ContactAuthVerifier interface {
	VerifyContactAuthToken(
		ctx context.Context,
		provider string,
		token string,
	) (*ContactAuthClaims, error)
}
