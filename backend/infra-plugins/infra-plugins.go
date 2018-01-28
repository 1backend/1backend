package infrapack

import (
	"github.com/1backend/1backend/backend/domain"
	"github.com/1backend/1backend/backend/infra-plugins/mysql"
	infrat "github.com/1backend/1backend/backend/infra-plugins/types"
)

// Plugins returns the plugin the project uses based on its mode and recipe
func Plugins(proj *domain.Project) []infrat.Plugin {
	ret := []infrat.Plugin{}
	for _, dep := range proj.Dependencies {
		switch dep.Type {
		case "mysql":
			ret = append(ret, mysqlpack.NewPack(proj))
		}
	}
	return ret
}
