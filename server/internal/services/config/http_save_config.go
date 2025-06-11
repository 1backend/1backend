/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package configservice

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	config "github.com/1backend/1backend/server/internal/services/config/types"
	types "github.com/1backend/1backend/server/internal/services/config/types"
	"github.com/pkg/errors"
)

// @Id saveConfig
// @Summary Save Config
// @Description Save the provided configuration to the server
// @Tags Config Svc
// @Accept json
// @Produce json
// @Param body body config.SaveConfigRequest true "Save Config Request"
// @Success 200 {object} config.SaveConfigResponse "Save Config Response"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /config-svc/config [put]
func (cs *ConfigService) Save(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, statusCode, err := cs.options.PermissionChecker.HasPermission(
		r,
		config.PermissionConfigEdit,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.GetAuthorized() {
		endpoint.Unauthorized(w)
		return
	}

	req := &config.SaveConfigRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		logger.Error("Failed to decode request body", slog.Any("error", err))
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	isAdmin, err := cs.options.Authorizer.IsAdminFromRequest(cs.publicKey, r)
	if err != nil {
		logger.Error("Failed to check admin status", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	err = cs.saveConfig(
		*isAuthRsp.App,
		isAdmin,
		isAuthRsp.User.Slug,
		req,
	)
	if err != nil {
		logger.Error("Failed to save config", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	jsonData, _ := json.Marshal(config.SaveConfigResponse{})
	_, err = w.Write([]byte(jsonData))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func (cs *ConfigService) saveConfig(
	app string,
	isAdmin bool,
	callerSlug string,
	newConfig *types.SaveConfigRequest,
) error {
	callerSlug = kebabToCamel(callerSlug)

	cs.configMutex.Lock()
	defer cs.configMutex.Unlock()

	newConfig.App = app

	oldConfigData := cs.configs[app]
	if oldConfigData == nil {
		oldConfigData = map[string]interface{}{}
		cs.configs[app] = oldConfigData
	}

	if isAdmin {
		oldConfigData = newConfig.Data
	} else {
		oldConfigData[callerSlug] = newConfig.Data[callerSlug]
	}

	newConfig.Data = oldConfigData

	newJson, err := json.Marshal(newConfig.Data)
	if err != nil {
		return errors.Wrap(err, "failed to marshal config")
	}
	newConfig.DataJSON = string(newJson)
	newConfig.Data = nil

	err = cs.configStore.Upsert(&types.Config{
		Id:       newConfig.Id,
		App:      newConfig.App,
		DataJSON: newConfig.DataJSON,
		Data:     newConfig.Data,
	})
	if err != nil {
		return errors.Wrap(err, "failed to save config")
	}

	ev := types.EventConfigUpdate{}
	_, err = cs.options.ClientFactory.Client(client.WithToken(cs.token)).
		FirehoseSvcAPI.PublishEvent(context.Background()).
		Event(openapi.FirehoseSvcEventPublishRequest{
			Event: &openapi.FirehoseSvcEvent{
				Name: openapi.PtrString(ev.Name()),
				Data: nil,
			},
		}).
		Execute()

	if err != nil {
		logger.Error("Failed to publish firehose event", slog.Any("error", err))
	}

	return nil
}

func kebabToCamel(s string) string {
	parts := strings.Split(s, "-")
	for i := 1; i < len(parts); i++ {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
		}
	}
	return strings.Join(parts, "")
}
