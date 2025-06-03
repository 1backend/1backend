/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package registryservice

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	registry "github.com/1backend/1backend/server/internal/services/registry/types"
	"github.com/pkg/errors"
)

// @ID registerInstance
// @Summary Register Instance
// @Description Registers an instance. Idempotent.
// @Tags Registry Svc
// @Accept json
// @Produce json
// @Param body body registry.RegisterInstanceRequest true "Register Instance Request"
// @Success 201 {object} registry.RegisterInstanceResponse
// @Failure 400 {object} registry.ErrorResponse "Invalid JSON"
// @Failure 401 {object} registry.ErrorResponse "Unauthorized"
// @Failure 500 {object} registry.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /registry-svc/instance [put]
func (rs *RegistryService) RegisterInstance(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, statusCode, err := rs.options.PermissionChecker.HasPermission(
		r,
		registry.PermissionInstanceEdit,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.GetAuthorized() {
		endpoint.Unauthorized(w)
		return
	}

	req := &registry.RegisterInstanceRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		logger.Error(
			"Failed to decode request",
			slog.Any("error", err),
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	err = rs.registerInstance(isAuthRsp.User.Slug, req)
	if err != nil {
		logger.Error(
			"Error registering instance",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	go func() {
		rs.triggerChan <- struct{}{}
		// @todo remove this and the tests fail ???
		rs.triggerChan <- struct{}{}
	}()

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(`{}`))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func (rs *RegistryService) registerInstance(
	callerSlug string,
	req *registry.RegisterInstanceRequest,
) error {
	var instance registry.Instance

	instances, err := rs.instanceStore.Query(datastore.Equals([]string{"url"}, req.URL)).
		Find()
	if err != nil {
		return err
	}

	if len(instances) > 0 {
		instance = *instances[0].(*registry.Instance)
	} else {
		instance.Id = sdk.Id("inst")
	}

	nu := len(instances) == 0
	if nu {
		instance.Slug = callerSlug
	}

	if callerSlug != "deploy-svc" && callerSlug != instance.Slug {
		return fmt.Errorf(
			"caller slug '%s' does not match instance slug '%s'. perhaps a misconfigured service already took this URL. if you are an admin you can fix that by calling 'oo instance remove %v'",
			callerSlug,
			instance.Slug,
			instance.Id,
		)
	}

	if req.Id != "" {
		instance.Id = req.Id
	}
	if req.URL != "" {
		instance.URL = req.URL
	}
	if req.Scheme != "" {
		instance.Scheme = req.Scheme
	}
	if req.Host != "" {
		instance.Host = req.Host
	}
	if req.IP != "" {
		instance.IP = req.IP
	}
	if req.Path != "" {
		instance.Path = req.Path
	}
	if req.DeploymentId != "" {
		instance.DeploymentId = req.DeploymentId
	}

	if instance.Status == "" {
		instance.Status = registry.InstanceStatusUnknown
	}

	if instance.URL == "" {
		return errors.New("url is missing")
	}

	return rs.instanceStore.Upsert(&instance)
}
