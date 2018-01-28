package techplugins

import (
	"fmt"

	"github.com/1backend/1backend/backend/domain"
	"github.com/1backend/1backend/backend/tech-plugins/go"
	"github.com/1backend/1backend/backend/tech-plugins/nodejs"
	techt "github.com/1backend/1backend/backend/tech-plugins/types"
	"github.com/1backend/1backend/backend/tech-plugins/typescript"
)

// GetProvider returns the plugin the project uses based on its mode and recipe
func Plugin(proj *domain.Project) techt.Plugin {
	switch {
	case proj.Mode == "go" && proj.Recipe == "":
		return goplugin.NewPlugin(proj)
	case proj.Mode == "nodejs" && proj.Recipe == "":
		return nodejsplugin.NewPlugin(proj)
	case proj.Mode == "typescript" && proj.Recipe == "":
		return typescriptplugin.NewPlugin(proj)
	}
	panic(fmt.Sprintf("Can't find recipe for mode %v and recipe %v", proj.Mode, proj.Recipe))
}
