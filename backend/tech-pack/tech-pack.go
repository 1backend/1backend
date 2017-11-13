package techpack

import (
	"fmt"
	"text/template"

	"github.com/1backend/1backend/backend/domain"
	"github.com/1backend/1backend/backend/tech-pack/go"
	"github.com/1backend/1backend/backend/tech-pack/nodejs"
	"github.com/1backend/1backend/backend/tech-pack/typescript"
)

// GetProvider returns the plugin the project uses based on its mode and recipe
func GetProvider(proj *domain.Project) Provider {
	switch {
	case proj.Mode == "go" && proj.Recipe == "":
		return gopack.NewPack(proj)
	case proj.Mode == "nodejs" && proj.Recipe == "":
		return nodejspack.NewPack(proj)
	case proj.Mode == "typescript" && proj.Recipe == "":
		return typescriptpack.NewPack(proj)
	}
	panic(fmt.Sprintf("Can't find recipe for mode %v and recipe %v", proj.Mode, proj.Recipe))
}

// Provider is the interface every plugin must implement
type Provider interface {
	// Folder name of tech pack
	RecipePath() string
	// Used to set up default endpoints and such
	CreateProjectPlugin() error
	// Returns filename of the file generated from code.tpl
	Outfile() string
	AddTemplateFuncs(fm *template.FuncMap)
	// FilesToBuild returns a list of files that one should write to the build folder
	// Returns a list of []string{file name, file content}
	FilesToBuild() [][]string
}
