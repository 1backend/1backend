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
	"net/http"

	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/datastore"
	deploy "github.com/openorch/openorch/server/internal/services/deploy/types"
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

	isAuthRsp, _, err := ns.clientFactory.Client(sdk.WithTokenFromRequest(r)).
		UserSvcAPI.IsAuthorized(r.Context(), *deploy.PermissionDeploymentView.Id).
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

	req := deploy.DeleteDeploymentRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	err = ns.deleteDeployment(req.DeploymentId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	go func() {
		ns.triggerChan <- struct{}{}
	}()

	w.Write([]byte(`{}`))
}

// This method is designed to be called frm the API endpoint, if you want to modify internal
// fields do a direct update.
func (ns *DeployService) deleteDeployment(deploymentId string) error {
	return ns.deploymentStore.Query(
		datastore.Equals(datastore.Field("id"), deploymentId),
	).Delete()
}
