package apipack

import (
	"github.com/1backend/1backend/backend/api-pack/go"
	apiTypes "github.com/1backend/1backend/backend/api-pack/types"
	"github.com/1backend/1backend/backend/domain"
)

// GetGenerators returns all the generators to
func Generators(proj *domain.Project) []Generator {
	return []Generator{
		goapi.NewGenerator(proj),
	}
}

type Generator interface {
	FilesToBuild(apiTypes.Context) ([][]string, error)
	// The subfolder to write files into eg. go, node, ts
	FolderName() string
}
