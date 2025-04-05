module github.com/1backend/1backend/examples/go/services/basic

go 1.23.0

toolchain go1.23.2

replace (
	github.com/1backend/1backend/clients/go => ../../../../clients/go/
	github.com/1backend/1backend/examples/go/services/basic/client/go => ./client/go
	github.com/1backend/1backend/sdk/go => ../../../../sdk/go/
)

require (
	github.com/1backend/1backend/clients/go v0.0.0-20250315124453-e49718380a69
	github.com/1backend/1backend/examples/go/services/basic/client/go v0.0.0-00010101000000-000000000000
	github.com/1backend/1backend/sdk/go v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.1
	github.com/pkg/errors v0.9.1
	github.com/samber/lo v1.49.1
	github.com/stretchr/testify v1.10.0
	github.com/swaggo/swag v1.16.4
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/andybalholm/brotli v1.1.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/flusflas/dipper v0.2.1 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/spec v0.20.4 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-sql-driver/mysql v1.9.1 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/sony/sonyflake v1.2.0 // indirect
	go.uber.org/mock v0.5.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	golang.org/x/tools v0.22.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
