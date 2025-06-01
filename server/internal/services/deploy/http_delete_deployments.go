/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package deployservice

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	deploy "github.com/1backend/1backend/server/internal/services/deploy/types"
)

// @ID deleteDeployment
// @Summary Delete Deployment
// @Description Delete a deployment.
// @Tags Deploy Svc
// @Accept json
// @Produce json
// @Param body body deploy.DeleteDeploymentRequest false "Delete Deploys Request"
// @Success 200 {object} deploy.DeleteDeploymentResponse
// @Failure 400 {object} deploy.ErrorResponse "Invalid JSON"
// @Failure 401 {object} deploy.ErrorResponse "Unauthorized"
// @Failure 500 {object} deploy.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /deploy-svc/deployment [delete]
func (ns *DeployService) DeleteDeployment(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, statusCode, err := ns.options.PermissionChecker.HasPermission(
		r,
		deploy.PermissionDeploymentView,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.GetAuthorized() {
		endpoint.Unauthorized(w)
		return
	}

	req := deploy.DeleteDeploymentRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error(
			"Error decoding request",
			slog.Any("error", err),
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	err = ns.deleteDeployment(req.DeploymentId)
	if err != nil {
		logger.Error(
			"Error deleting deployment",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	go func() {
		ns.triggerChan <- struct{}{}
	}()

	_, err = w.Write([]byte(`{}`))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

// This method is designed to be called frm the API endpoint, if you want to modify internal
// fields do a direct update.
func (ns *DeployService) deleteDeployment(deploymentId string) error {
	return ns.deploymentStore.Query(
		datastore.Equals(datastore.Field("id"), deploymentId),
	).Delete()
}
