package techpack

import (
	"fmt"

	"github.com/1backend/1backend/backend/api-packs/go"
	apiTypes "github.com/1backend/1backend/backend/api-packs/types"
	"github.com/1backend/1backend/backend/domain"
)

// GetProvider returns the plugin the project uses based on its mode and recipe
func GetGenerator(proj *domain.Project, language string) Generator {
	switch {
	case language == "go":
		return goapi.NewGenerator(proj)
	}
	panic(fmt.Sprintf("Can't find generator %v for language", language))
}

type Generator interface {
	GenerateApi(apiTypes.Context) (string, error)
}
