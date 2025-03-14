/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package configservice

import (
	"encoding/json"
	"net/http"
	"path"

	config "github.com/openorch/openorch/server/internal/services/config/types"
	types "github.com/openorch/openorch/server/internal/services/config/types"
	modelservice "github.com/openorch/openorch/server/internal/services/model"
	"github.com/pkg/errors"
)

// @Id getConfig
// @Summary Get Config
// @Description Fetch the current configuration from the server
// @Tags Config Svc
// @Accept json
// @Produce json
// @Param namespace query string false "Namespace"
// @Success 200 {object} config.GetConfigResponse "Current configuration retrieved successfully"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /config-svc/config [get]
func (cs *ConfigService) Get(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Config get should not be authorized because it is public, nonsensitive information.
	// Think about app config, A/B tests and such.

	q := r.URL.Query()
	namespace := q.Get("namespace")
	if namespace == "" {
		namespace = "default"
	}

	conf, err := cs.getConfig(namespace)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(config.GetConfigResponse{
		Config: conf,
	})
	w.Write(jsonData)
}

func (cs *ConfigService) getConfig(namespace string) (*types.Config, error) {
	if namespace == "" {
		namespace = "default"
	}
	data, ok := cs.configs[namespace]
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
				cs.homeDir,
				"downloads",
			),
		}
	}

	if ret.Data["model-svc"] == nil {
		ret.Data["model-svc"] = map[string]interface{}{
			"currentModelId": modelservice.DefaultModelId,
		}
	}

	if ret.Data["config-svc"] == nil {
		ret.Data["config-svc"] = map[string]interface{}{
			"directory": cs.homeDir,
		}
	}
}
