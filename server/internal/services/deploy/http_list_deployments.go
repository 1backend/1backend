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

	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	deploy "github.com/1backend/1backend/server/internal/services/deploy/types"
)

// @ID listDeployments
// @Summary List Deployments
// @Description Retrieve a list of deployments.
// @Tags Deploy Svc
// @Accept json
// @Produce json
// @Param body body deploy.ListDeploymentsRequest false "List Deploys Request"
// @Success 200 {object} deploy.ListDeploymentsResponse
// @Failure 400 {object} deploy.ErrorResponse "Invalid JSON"
// @Failure 401 {object} deploy.ErrorResponse "Unauthorized"
// @Failure 500 {object} deploy.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /deploy-svc/deployments [post]
func (ns *DeployService) ListDeployments(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, statusCode, err := ns.permissionChecker.HasPermission(
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

	req := deploy.ListDeploymentsRequest{}
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

	deployments, err := ns.listDeployments()
	if err != nil {
		logger.Error(
			"Error listing deployments",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	response := deploy.ListDeploymentsResponse{
		Deployments: deployments,
	}

	bs, _ := json.Marshal(response)
	_, err = w.Write(bs)
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func (ns *DeployService) listDeployments() ([]*deploy.Deployment, error) {
	deploymentIs, err := ns.deploymentStore.Query().Find()
	if err != nil {
		return nil, err
	}

	ret := []*deploy.Deployment{}
	for _, deploymentI := range deploymentIs {
		ret = append(ret, deploymentI.(*deploy.Deployment))
	}

	return ret, err
}
