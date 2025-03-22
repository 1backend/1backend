module first-service-go

go 1.23.0

replace github.com/1backend/1backend/clients/go => ../../../../clients/go

replace github.com/1backend/1backend/sdk/go => ../../../../sdk/go

require (
	github.com/1backend/1backend/clients/go v0.0.0-20250315124453-e49718380a69
	github.com/1backend/1backend/sdk/go v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.9.1
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/flusflas/dipper v0.2.1 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/sony/sonyflake v1.2.0 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	go.uber.org/mock v0.5.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
