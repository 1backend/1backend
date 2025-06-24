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
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	config "github.com/1backend/1backend/server/internal/services/config/types"
	types "github.com/1backend/1backend/server/internal/services/config/types"
	"github.com/pkg/errors"
)

// @Id saveConfig
// @Summary Save Config
// @Description Save the provided configuration to the server.
// @Description The app from the caller's token is used to determine which app the config belongs to.
// @Description The caller's camelCased slug (e.g., "test-user-slug" becomes "testUserSlug") is used as the config key automatically,
// @Description except for users who have the "config-svc:config:edit-on-behalf" permission (admins), who can specify any key they want.
// @Description Admins (users with the "config-svc:config:edit-on-behalf" permission) can also provide an "app" field in the request body to specify which app the config belongs to, while
// @Description non-admin users cannot specify the "app" field, the app associated with their token will be used.
// @Description
// @Description The save performs a deep merge, that is:
// @Description - Nested objects are recursively merged rather than replaced.
// @Description - If a field exists in both the existing and the incoming config and both values are objects, their contents are merged.
// @Description - If a field exists in both but one or both values are not objects (e.g., string, number, array), the incoming value replaces the existing one.
// @Description - Fields present only in the incoming config are added.
// @Description - Fields present only in the existing config are preserved.
// @Description - Top-level and nested merges follow the same rules.
// @Tags Config Svc
// @Accept json
// @Produce json
// @Param body body config.SaveConfigRequest true "Save Config Request"
// @Success 200 {object} config.SaveConfigResponse "Save Config Response"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /config-svc/config [put]
func (cs *ConfigService) SaveConfig(
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

	canActonBehalf := false
	isAuthRsp, _, err = cs.options.PermissionChecker.HasPermission(
		r,
		config.PermissionConfigEditOnBehalf,
	)

	if err == nil && isAuthRsp.GetAuthorized() {
		canActonBehalf = true
	}

	if !canActonBehalf && req.App != "" {
		logger.Error("Unauthorized attempt to save config with app specified",
			slog.String("app", req.App),
		)
		endpoint.Unauthorized(w)
		return
	}

	app := req.App
	if app == "" {
		app = *isAuthRsp.App
	}

	err = cs.saveConfig(
		canActonBehalf,
		app,
		isAuthRsp.User.Slug,
		req,
	)
	if err != nil {
		logger.Error("Failed to save config",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	endpoint.WriteJSON(w, http.StatusOK, config.SaveConfigResponse{})
}

func (cs *ConfigService) saveConfig(
	canActonBehalf bool,
	app string,
	callerSlug string,
	newConfig *types.SaveConfigRequest,
) error {
	callerSlug = kebabToCamel(callerSlug)

	if canActonBehalf {
		if newConfig.Key == "" {
			return errors.New("key must be provided when editing on behalf of another user")
		}
		callerSlug = kebabToCamel(newConfig.Key)
	}

	cs.configMutex.Lock()
	defer cs.configMutex.Unlock()

	now := time.Now()

	id := fmt.Sprintf("%s_%s", app, callerSlug)

	existing, found, err := cs.configStore.Query(
		datastore.Id(id),
	).FindOne()
	if err != nil {
		return errors.Wrap(err, "failed to query config")
	}

	entry := &types.Config{}

	if !found {
		entry.Id = id
		entry.Data = map[string]interface{}{}
		entry.DataJSON = "{}"
		entry.CreatedAt = now
		entry.UpdatedAt = now
	} else {
		entry = existing.(*types.Config)
		err = json.Unmarshal([]byte(entry.DataJSON), &entry.Data)
		if err != nil {
			return errors.Wrap(err, "failed to unmarshal existing config data")
		}
	}

	entry.Data = DeepMerge(entry.Data, newConfig.Data)

	newJson, err := json.Marshal(entry.Data)
	if err != nil {
		return errors.Wrap(err, "failed to marshal config")
	}

	entry.DataJSON = string(newJson)
	entry.Data = nil

	err = cs.configStore.Upsert(entry)
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
