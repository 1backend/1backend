package infrapack

import (
	"text/template"

	"github.com/1backend/1backend/backend/domain"
	"github.com/1backend/1backend/backend/infra-pack/mysql"
)

// GetProvider returns the plugin the project uses based on its mode and recipe
func GetProviders(proj *domain.Project) []Provider {
	ret := []Providers{}
	for _, dep := range proj.Dependencies {
		switch dep.Name {
		case "mysql":
			ret = append(ret, mysql.NewPack(proj))
		}
	}
	return ret
}

// Provider is the interface every tech pack plugin must implement
type Provider interface {
	Name() string
	// Used to set up default endpoints and such
	CreateProjectPlugin() error
	// Func added here will be available when generated both the global and the readme section
	AddTemplateFuncs(fm *template.FuncMap)
	// String returned is appended to the global variables section of the code.tpl
	// being generated. Return appropriate code for the mode (language) of the project.
	GlobalsSection() (string, error)
	// String returned is appended to the readme.
	// Use this method to describe global variables set up and such
	// see https://github.com/1backend/1backend/issues/126
	ReadmeSection() (string, error)
	// This is called pre build when updating a project
	UpdateProjectPlugin() error
}
