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

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	config "github.com/1backend/1backend/server/internal/services/config/types"
	types "github.com/1backend/1backend/server/internal/services/config/types"
	modelservice "github.com/1backend/1backend/server/internal/services/model"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
)

// @Id listConfigs
// @Summary List Configs
// @Description Retrieves the current configurations for a specified app.
// @Description Since any user can save configurations, it is strongly advised that you supply a list of
// @Description owners to filter on.
// @Description If no app is specified, the default "unnamed" app is used.
// @Description This is a public endpoint and does not require authentication.
// @Description Configuration data is non-sensitive. For sensitive data, refer to the Secret Service.
// @Description
// @Description Configurations are used to control frontend behavior, A/B testing, feature flags, and other non-sensitive settings.
// @Tags Config Svc
// @Accept json
// @Produce json
// @Param body body config.ListConfigsRequest false "List Configs Request"
// @Success 200 {object} config.ListConfigsResponse "Current configuration"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /config-svc/configs [post]
func (cs *ConfigService) ListConfigs(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Config get should not be authorized because
	// it is public, nonsensitive information.

	req := &config.ListConfigsRequest{}

	if r.ContentLength > 0 {
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			logger.Error("Failed to decode request body", slog.Any("error", err))
			endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
			return
		}
		defer r.Body.Close()
	}

	if req.AppHost == "" {
		var err error
		req.AppHost, err = cs.options.TokenExchanger.AppHostFromRequest(r)
		if err != nil {
			logger.Error(
				"Failed to get app from request host",
				slog.Any("error", err),
			)
			endpoint.InternalServerError(w)
			return
		}
	}

	appId, err := cs.options.TokenExchanger.AppIdFromHost(
		r.Context(),
		req.AppHost,
	)
	if err != nil {
		logger.Error(
			"Failed to get app id from host",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	confs, err := cs.listConfigs(appId, req)
	if err != nil {
		logger.Error("Failed to get config", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	jsonData, _ := json.Marshal(config.ListConfigsResponse{
		Configs: confs,
	})
	_, err = w.Write([]byte(jsonData))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func (cs *ConfigService) listConfigs(
	appId string,
	req *types.ListConfigsRequest,
) (map[string]*types.Config, error) {
	ids := []any{}
	for _, slug := range req.Ids {
		slug = kebabToCamel(slug)
		ids = append(ids,
			slug,
		)
	}
	for slug := range req.Selector {
		slug = kebabToCamel(slug)
		ids = append(ids,
			slug,
		)
	}

	filters := []datastore.Filter{}

	if len(ids) > 0 {
		filters = append(filters, datastore.IsInList(
			datastore.Field("id"),
			ids...,
		))
	}

	if req.AppHost != "" {

		filters = append(filters, datastore.Equals(
			datastore.Field("appId"),
			appId,
		))
	}

	configIs, err := cs.configStore.Query(
		filters...,
	).Find()
	if err != nil {
		return nil, errors.Wrap(err, "failed to query configs")
	}

	ret := map[string]*types.Config{}

	for _, configI := range configIs {
		conf := configI.(*types.Config)
		conf.InternalId = ""

		if conf.DataJSON != "" {
			if req.Selector != nil && req.Selector[conf.Id] != nil {
				conf.Data = map[string]any{}
				for _, path := range req.Selector[conf.Id] {
					res := gjson.Get(conf.DataJSON, path)
					if res.Exists() {
						m := ExpandDotPath(path, res.Value())
						conf.Data = DeepMerge(conf.Data, m)
					}
				}
			} else {
				err = json.Unmarshal([]byte(conf.DataJSON), &conf.Data)
				if err != nil {
					logger.Error("Failed to unmarshal config data",
						slog.Any("error", err),
						slog.String("dataJson", conf.DataJSON),
					)

					return nil, errors.Wrap(err, "failed to unmarshal config")
				}
			}
		}

		conf.DataJSON = ""
		ret[conf.Id] = conf
	}

	cs.defaults(req, ret)

	return ret, nil
}

func (cs *ConfigService) defaults(
	req *types.ListConfigsRequest,
	ret map[string]*types.Config,
) {
	slugMap := map[string]bool{}
	for _, slug := range req.Ids {
		slug = kebabToCamel(slug)
		slugMap[slug] = true
	}

	if slugMap["fileSvc"] && ret["fileSvc"] == nil {
		ret["fileSvc"] = &types.Config{
			Data: map[string]interface{}{
				"downloadFolder": path.Join(
					cs.options.HomeDir,
					"downloads",
				),
			},
		}
	}

	if slugMap["modelSvc"] && ret["modelSvc"] == nil {
		ret["modelSvc"] = &types.Config{
			Data: map[string]interface{}{
				"currentModelId": modelservice.DefaultModelId,
			},
		}
	}

	if slugMap["configSvc"] && ret["configSvc"] == nil {
		ret["configSvc"] = &types.Config{
			Data: map[string]interface{}{
				"directory": cs.options.HomeDir,
			},
		}
	}
}
