/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package containerservice

import (
	"encoding/json"
	"net/http"

	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/datastore"
	container "github.com/openorch/openorch/server/internal/services/container/types"
)

// @ID listContainerLogs
// @Summary List Logs
// @Description List Container logs.
// @Description
// @Description Requires the `container-svc:log:view` permission.
// @Tags Container Svc
// @Accept json
// @Produce json
// @Param body body container.ListLogsRequest true "List Logs Request"
// @Success 200 {object} container.ListLogsResponse
// @Failure 400 {object} container.ErrorResponse "Invalid JSON"
// @Failure 401 {object} container.ErrorResponse "Unauthorized"
// @Failure 500 {object} container.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /container-svc/logs [post]
func (dm *ContainerService) ListLogs(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, _, err := dm.clientFactory.Client(sdk.WithTokenFromRequest(r)).
		UserSvcAPI.IsAuthorized(r.Context(), *container.PermissionContainerCreate.Id).
		Body(openapi.UserSvcIsAuthorizedRequest{
			GrantedSlugs: []string{"model-svc", "deploy-svc"},
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

	req := &container.ListLogsRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	rsp, err := dm.listLogs(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(rsp)
	w.Write(jsonData)
}

func (dm *ContainerService) listLogs(req *container.ListLogsRequest) (*container.ListLogsResponse, error) {
	q := dm.logStore.Query()

	if req.NodeId != "" {
		q = dm.logStore.Query(datastore.Equals([]string{"nodeId"}, req.NodeId))
	}

	if req.ContainerId != "" {
		q = dm.logStore.Query(datastore.Equals([]string{"containerId"}, req.ContainerId))
	}

	q = q.OrderBy(datastore.OrderByField("createdAt", true))

	if req.Limit != 0 {
		q.Limit(req.Limit)
	}

	resI, err := q.Find()
	if err != nil {
		return nil, err
	}

	res := []*container.Log{}
	for _, v := range resI {
		res = append(res, v.(*container.Log))
	}

	return &container.ListLogsResponse{
		Logs: res,
	}, nil
}
