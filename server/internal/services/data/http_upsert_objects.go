/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package dynamicservice

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	data "github.com/1backend/1backend/server/internal/services/data/types"
)

// @ID upsertObjects
// @Summary Upsert Objects
// @Description Upserts objects by ids.
// @Tags Data Svc
// @Accept json
// @Produce json
// @Param body body data.UpsertObjectRequest true "Upsert request payload"
// @Success 200 {object} data.UpsertObjectResponse "Successful upsert of objects"
// @Failure 400 {object} data.ErrorResponse "Invalid JSON"
// @Failure 401 {object} data.ErrorResponse "Unauthorized"
// @Failure 500 {object} data.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /data-svc/objects/upsert [put]
func (g *DataService) SaveObjects(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, statusCode, err := g.options.PermissionChecker.HasPermission(
		r,
		data.PermissionObjectCreate,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.GetAuthorized() {
		endpoint.Unauthorized(w)
		return
	}

	req := &data.UpsertObjectsRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		logger.Error(
			"Error decoding request body",
			slog.Any("error", err),
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	err = g.upsertObjects(req)
	if err != nil {
		logger.Error(
			"Error upserting objects",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	_, err = w.Write([]byte(`{}`))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func (g *DataService) upsertObjects(
	request *data.UpsertObjectsRequest,
) error {
	objectIs := []datastore.Row{}
	for _, object := range request.Objects {
		objectIs = append(objectIs, object)
	}
	return g.store.UpsertMany(objectIs)
}
