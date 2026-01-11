package universe

import (
	"net/http"
	"time"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/infra"
	"github.com/1backend/1backend/sdk/go/lock"
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

	// eg. mysql, postgres
	Db string

	// Connection string eg.
	// "postgres://postgres:mysecretpassword@localhost:5432/mydatabase?sslmode=disable"
	DbConnectionString string

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

	// URL of the local 1Backend server instance
	Url string

	// FileCacheItemMaxSize defines the maximum size of a single file allowed in the cache.
	// Use bytes (e.g., 50 * 1024 * 1024 for 50MB).
	FileCacheItemMaxSize int64

	// FileCacheMaxSize defines the total disk/memory footprint allowed for the file cache.
	FileCacheMaxSize int64

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
