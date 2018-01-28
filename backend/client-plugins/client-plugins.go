package apipack

import (
	goclient "github.com/1backend/1backend/backend/client-plugins/go"
	ngclient "github.com/1backend/1backend/backend/client-plugins/go"
	clientt "github.com/1backend/1backend/backend/client-plugins/types"
	"github.com/1backend/1backend/backend/domain"
)

// Plugins returns all the generators to
func Plugins(proj *domain.Project) []clientt.ClientPlugin {
	return []clientt.ClientPlugin{
		goclient.New(proj),
		ngclient.New(proj),
	}
}
