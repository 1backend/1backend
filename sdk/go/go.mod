module github.com/openorch/openorch/sdk/go

go 1.23

// replace github.com/openorch/openorch/clients/go => ../../clients/go

require (
	github.com/andybalholm/brotli v1.1.1
	github.com/davecgh/go-spew v1.1.1
	github.com/flusflas/dipper v0.2.1
	github.com/go-sql-driver/mysql v1.8.1
	github.com/golang-jwt/jwt/v4 v4.5.1
	github.com/google/uuid v1.6.0
	github.com/lib/pq v1.10.9
	github.com/openorch/openorch/clients/go v0.0.0-20241204104942-d2837b21586b
	github.com/pkg/errors v0.9.1
	github.com/sony/sonyflake v1.2.0
	github.com/stretchr/testify v1.10.0
	go.uber.org/mock v0.5.0
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
