/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package configservice

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"path"

	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	config "github.com/1backend/1backend/server/internal/services/config/types"
	types "github.com/1backend/1backend/server/internal/services/config/types"
	modelservice "github.com/1backend/1backend/server/internal/services/model"
	"github.com/pkg/errors"
)

// @Id readConfig
// @Summary Read Config
// @Description Retrieves the current configuration for a specified app.
// @Description If no app is specified, the default "unnamed" app is used.
// @Description This is a public endpoint and does not require authentication.
// @Description Configuration data is non-sensitive. For sensitive data, refer to the Secret Service.
// @Description
// @Description Configurations are used to control frontend behavior, A/B testing, feature flags, and other non-sensitive settings.
// @Tags Config Svc
// @Accept json
// @Produce json
// @Param app query string false "App"
// @Success 200 {object} config.GetConfigResponse "Current configuration"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /config-svc/config [get]
func (cs *ConfigService) Get(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Config get should not be authorized because it is public, nonsensitive information.
	// Think about app config, A/B tests and such.

	q := r.URL.Query()
	app := q.Get("app")
	if app == "" {
		app = defaultApp
	}

	conf, err := cs.readConfig(app)
	if err != nil {
		logger.Error("Failed to get config", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	jsonData, _ := json.Marshal(config.GetConfigResponse{
		Config: conf,
	})
	_, err = w.Write([]byte(jsonData))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func (cs *ConfigService) readConfig(app string) (*types.Config, error) {
	data, ok := cs.configs[app]
	if !ok {
		conf := &types.Config{
			Data: map[string]interface{}{},
		}
		cs.mod(conf)
		return conf, nil
	}
	v, err := json.Marshal(data)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal config")
	}

	ret := &types.Config{}
	err = json.Unmarshal(v, &ret.Data)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal config")
	}

	cs.mod(ret)

	return ret, nil
}

func (cs ConfigService) mod(ret *types.Config) {
	if ret.Data == nil {
		ret.Data = map[string]interface{}{}
	}

	if ret.Data["file-svc"] == nil {
		ret.Data["file-svc"] = map[string]interface{}{
			"downloadFolder": path.Join(
				cs.options.HomeDir,
				"downloads",
			),
		}
	}

	if ret.Data["modelSvc"] == nil {
		ret.Data["modelSvc"] = map[string]interface{}{
			"currentModelId": modelservice.DefaultModelId,
		}
	}

	if ret.Data["config-svc"] == nil {
		ret.Data["config-svc"] = map[string]interface{}{
			"directory": cs.options.HomeDir,
		}
	}
}
