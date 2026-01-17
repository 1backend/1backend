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
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	config "github.com/1backend/1backend/server/internal/services/config/types"
	types "github.com/1backend/1backend/server/internal/services/config/types"
	"github.com/pkg/errors"
)

var (
	idMustBeProvidedErr = errors.New("id must be provided when editing on behalf of another user")
	defaultBranch       = "main"
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
	if !isAuthRsp.Authorized {
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

	if err == nil && isAuthRsp.Authorized {
		canActonBehalf = true
	}

	if !canActonBehalf {
		if req.AppHost != "" {
			logger.Error("Unauthorized attempt to save config with app specified",
				slog.String("appHost", req.AppHost),
			)
			endpoint.Unauthorized(w)
			return
		}

		if req.Id != "" {
			logger.Error("Unauthorized attempt to save config with key specified",
				slog.String("id", req.Id),
			)
			endpoint.Unauthorized(w)
			return
		}
	}

	if req.AppHost == "" {
		req.AppHost = isAuthRsp.App.Host
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

	err = cs.saveConfig(
		canActonBehalf,
		appId,
		req.AppHost,
		isAuthRsp.User.Slug,
		req,
	)
	switch {
	case errors.Is(err, idMustBeProvidedErr):
		endpoint.WriteString(w, http.StatusBadRequest, err.Error())
		return
	case err != nil:
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
	appId string,
	appHost string,
	callerSlug string,
	newConfig *types.SaveConfigRequest,
) error {
	callerSlug = kebabToCamel(callerSlug)

	if canActonBehalf {
		if newConfig.Id == "" {
			return errors.New("id must be provided when editing on behalf of another user for safety reasons")
		}
		callerSlug = kebabToCamel(newConfig.Id)
	}

	cs.configMutex.Lock()
	defer cs.configMutex.Unlock()

	now := time.Now()

	id := callerSlug

	branch := defaultBranch
	if newConfig.Branch != "" {
		branch = newConfig.Branch
	}

	existing, found, err := cs.configStore.Query(
		datastore.Id(id),
		datastore.Equals([]string{"appId"}, appId),
		datastore.Equals([]string{"branch"}, branch),
	).FindOne()
	if err != nil {
		return errors.Wrap(err, "failed to query config")
	}

	if newConfig.Id == "" {
		newConfig.Id = callerSlug
	}

	entry := &types.Config{}

	if !found {
		entry.InternalId, err = sdk.InternalId(appId, sdk.DeterministicId(id, branch))
		if err != nil {
			return fmt.Errorf("failed to generate internal id: %w", err)
		}
		entry.Id = id
		entry.AppId = appId
		entry.Data = map[string]interface{}{}
		entry.DataJSON = "{}"
		entry.CreatedAt = now
		entry.UpdatedAt = now
		entry.Branch = branch
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

	versionId := sdk.OpaqueId("ver")
	versionInternalId, err := sdk.InternalId(appId, sdk.DeterministicId(id, versionId))
	if err != nil {
		return fmt.Errorf("failed to generate version internal id: %w", err)
	}

	version := config.Version{
		InternalId: versionInternalId,
		AppId:      appId,
		Id:         id,
		VersionId:  versionId,
		CreatedAt:  now,
		UpdatedAt:  now,
		DataJSON:   string(newJson),
		Branch:     branch,
	}

	err = cs.versionStore.Upsert(version)
	if err != nil {
		return errors.Wrap(err, "failed to save config version")
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

	cs.invalidate(appHost, branch)

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
