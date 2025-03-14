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
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/logger"
	config "github.com/openorch/openorch/server/internal/services/config/types"
	types "github.com/openorch/openorch/server/internal/services/config/types"
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
	isAuthRsp, _, err := cs.clientFactory.Client(sdk.WithTokenFromRequest(r)).
		UserSvcAPI.IsAuthorized(r.Context(), *config.PermissionConfigEdit.Id).
		Body(openapi.UserSvcIsAuthorizedRequest{
			GrantedSlugs: []string{"model-svc"},
		}).
		Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if !isAuthRsp.GetAuthorized() {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`Unauthorized`))
		return
	}

	req := &config.SaveConfigRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	isAdmin, err := cs.authorizer.IsAdminFromRequest(cs.publicKey, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = cs.saveConfig(isAdmin, *isAuthRsp.User.Slug, *req.Config)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(config.SaveConfigResponse{})
	w.Write(jsonData)
}

func (cs *ConfigService) saveConfig(isAdmin bool, callerSlug string, newConfig types.Config) error {
	if newConfig.Namespace == "" {
		newConfig.Namespace = "default"
	}

	cs.configMutex.Lock()
	defer cs.configMutex.Unlock()

	oldConfigData := cs.configs[newConfig.Namespace]
	if oldConfigData == nil {
		oldConfigData = map[string]interface{}{}
		cs.configs[newConfig.Namespace] = oldConfigData
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

	err = cs.configStore.Upsert(newConfig)
	if err != nil {
		return errors.Wrap(err, "failed to save config")
	}

	ev := types.EventConfigUpdate{}
	_, err = cs.clientFactory.Client(sdk.WithToken(cs.token)).
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
