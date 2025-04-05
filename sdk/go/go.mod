module github.com/1backend/1backend/sdk/go

go 1.23

// replace github.com/1backend/1backend/clients/go => ../../clients/go

require (
	github.com/1backend/1backend/clients/go v0.0.0-20250315124453-e49718380a69
	github.com/andybalholm/brotli v1.1.1
	github.com/davecgh/go-spew v1.1.1
	github.com/flusflas/dipper v0.2.1
	github.com/go-sql-driver/mysql v1.8.1
	github.com/golang-jwt/jwt/v5 v5.2.2
	github.com/google/uuid v1.6.0
	github.com/lib/pq v1.10.9
	github.com/pkg/errors v0.9.1
	github.com/samber/lo v1.49.1
	github.com/sony/sonyflake v1.2.0
	github.com/stretchr/testify v1.10.0
	go.uber.org/mock v0.5.0
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
