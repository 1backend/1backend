package types

// Provider is the interface every tech pack plugin must implement
type Plugin interface {
	// Pretty name of the infrastructure plugin ie. PostgreSQL
	Name() string
	// Func added here will be available when generated both the global and the readme section
	PreDeploy(envars map[string]string) (*PreDeploy, error)
}

// Output of predeploy data
type PreDeploy struct {
	// String returned is appended to the readme.
	// Use this method to describe available environment variables and such
	ReadmeSection string
	// Script output
	Output string
}
