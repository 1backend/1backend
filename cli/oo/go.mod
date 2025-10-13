module github.com/1backend/1backend/cli/oo

go 1.23

// replace github.com/1backend/1backend/clients/go => ../../clients/go

//replace github.com/1backend/1backend/sdk/go => ../../sdk/go

require (
	github.com/1backend/1backend/clients/go v0.0.0-20251013205843-dddc96a2a59c
	github.com/1backend/1backend/sdk/go v0.0.0-20251013205843-dddc96a2a59c
	github.com/fatih/color v1.18.0
	github.com/ghodss/yaml v1.0.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.8.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.11.1
	golang.org/x/term v0.27.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/andybalholm/brotli v1.1.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgraph-io/ristretto v0.2.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/samber/lo v1.49.1 // indirect
	github.com/sony/sonyflake v1.2.0 // indirect
	go.uber.org/mock v0.5.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
