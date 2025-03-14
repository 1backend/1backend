/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package registryservice

import (
	"encoding/json"
	"net/http"

	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	registry "github.com/openorch/openorch/server/internal/services/registry/types"
)

// @ID listNodes
// @Summary List Nodes
// @Description Retrieve a list of nodes.
// @Tags Registry Svc
// @Accept json
// @Produce json
// @Param body body registry.ListNodesRequest false "List Nodes Request"
// @Success 200 {object} registry.ListNodesResponse
// @Failure 400 {object} registry.ErrorResponse "Invalid JSON"
// @Failure 401 {object} registry.ErrorResponse "Unauthorized"
// @Failure 500 {object} registry.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /registry-svc/nodes [post]
func (ns *RegistryService) ListNodes(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, _, err := ns.clientFactory.Client(sdk.WithTokenFromRequest(r)).
		UserSvcAPI.IsAuthorized(r.Context(), *registry.PermissionNodeView.Id).
		Body(openapi.UserSvcIsAuthorizedRequest{
			GrantedSlugs: []string{
				"deploy-svc",
				"file-svc",
			},
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

	req := &registry.ListNodesRequest{}
	if r.ContentLength != 0 {
		err = json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`Invalid JSON`))
			return
		}
		defer r.Body.Close()
	}

	nodes, err := ns.listNodes(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response := registry.ListNodesResponse{
		Nodes: nodes,
	}

	bs, _ := json.Marshal(response)
	w.Write(bs)
}

func (ns *RegistryService) listNodes(req *registry.ListNodesRequest) ([]*registry.Node, error) {
	nodeIs, err := ns.nodeStore.Query().Find()
	if err != nil {
		return nil, err
	}

	ret := []*registry.Node{}

	for _, nodeI := range nodeIs {
		node := nodeI.(*registry.Node)

		match := len(req.Ids) == 0

		if len(req.Ids) != 0 {
			for _, id := range req.Ids {
				if id == node.Id {
					match = true
				}
			}
		}

		if match {
			ret = append(ret, node)
		}
	}

	return ret, err
}
