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

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	registry "github.com/1backend/1backend/server/internal/services/registry/types"
)

// @ID selfNode
// @Summary View Self Node
// @Description Show the local node.
// @Tags Registry Svc
// @Accept json
// @Produce json
// @Param body body registry.NodeSelfRequest false "List Registrys Request"
// @Success 200 {object} registry.NodeSelfResponse
// @Failure 400 {object} registry.ErrorResponse "Invalid JSON"
// @Failure 401 {object} registry.ErrorResponse "Unauthorized"
// @Failure 500 {object} registry.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /registry-svc/node/self [get]
func (ns *RegistryService) NodeSelf(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, statusCode, err := ns.options.PermissionChecker.HasPermission(
		r,
		registry.PermissionNodeView,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.Authorized {
		endpoint.Unauthorized(w)
		return
	}

	node, err := ns.thisNode()
	if err != nil {
		logger.Error(
			"Cannot find node",
			slog.String("error", err.Error()),
		)
		endpoint.InternalServerError(w)
		return
	}

	response := registry.NodeSelfResponse{
		Node: *node,
	}

	bs, _ := json.Marshal(response)
	_, err = w.Write(bs)
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func (ns *RegistryService) thisNode() (*registry.Node, error) {
	nodeIs, err := ns.nodeStore.Query(
		datastore.Equals([]string{"url"}, ns.options.Url),
	).Find()
	if err != nil {
		return nil, err
	}

	if len(nodeIs) == 0 {
		return nil, fmt.Errorf("cannot find node with url '%v'", ns.options.Url)
	}

	ret := nodeIs[0].(*registry.Node)

	return ret, err
}
