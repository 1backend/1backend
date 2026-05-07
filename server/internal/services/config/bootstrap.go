package configservice

import (
	"context"

	config "github.com/1backend/1backend/server/internal/services/config/types"
	"github.com/pkg/errors"
)

func (cs *ConfigService) BootstrapSaveConfigs(
	ctx context.Context,
	configs []config.SaveConfigRequest,
	appIDFromHost func(string) (string, error),
) error {
	if len(configs) == 0 {
		return nil
	}
	if appIDFromHost == nil {
		return errors.New("bootstrap app resolver is nil")
	}

	for i := range configs {
		req := configs[i]
		appHost := req.AppHost
		if appHost == "" {
			appHost = defaultApp
		}

		appID, err := appIDFromHost(appHost)
		if err != nil {
			return errors.Wrapf(err, "bootstrap app %s", appHost)
		}

		if err := cs.saveConfig(true, appID, appHost, "1backend", &req); err != nil {
			return errors.Wrapf(err, "bootstrap config %s for app %s", req.Id, appHost)
		}
	}

	return nil
}
