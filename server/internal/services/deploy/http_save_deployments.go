/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package deployservice

import (
	"encoding/json"
	"log/slog"
	"net/http"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	deploy "github.com/1backend/1backend/server/internal/services/deploy/types"
)

// @ID saveDeployment
// @Summary Save Deployment
// @Description Save a deployment.
// @Tags Deploy Svc
// @Accept json
// @Produce json
// @Param body body deploy.SaveDeploymentRequest false "Save Deploys Request"
// @Success 200 {object} deploy.SaveDeploymentResponse
// @Failure 400 {object} deploy.ErrorResponse "Invalid JSON"
// @Failure 401 {object} deploy.ErrorResponse "Unauthorized"
// @Failure 500 {object} deploy.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /deploy-svc/deployment [put]
func (ns *DeployService) SaveDeployment(
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

	req := deploy.SaveDeploymentRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error(
			"Error decoding request body",
			slog.Any("error", err),
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	err = ns.saveDeployment(req.Deployment)
	if err != nil {
		logger.Error(
			"Error saving deployment",
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
func (ns *DeployService) saveDeployment(deployment *deploy.Deployment) error {
	deployment.Status = deploy.DeploymentStatusPending

	if deployment.Replicas == 0 {
		deployment.Replicas = 1
	}

	if deployment.Id == "" {
		deployment.Id = sdk.Id("depl")
	} else {
		depIs, err := ns.deploymentStore.Query(datastore.Id(deployment.Id)).Find()
		if err != nil {
			return err
		}
		if len(depIs) == 0 {
			return ns.deploymentStore.Upsert(deployment)
		}

		oldDeployment := depIs[0].(*deploy.Deployment)

		deployment.Status = oldDeployment.Status
		deployment.Details = oldDeployment.Details

		return ns.deploymentStore.Upsert(deployment)
	}

	return ns.deploymentStore.Upsert(deployment)
}
