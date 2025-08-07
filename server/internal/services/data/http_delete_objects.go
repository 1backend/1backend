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
	"errors"
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	data "github.com/1backend/1backend/server/internal/services/data/types"
)

// @ID deleteObjects
// @Summary     Delete Objects
// @Description Deletes all objects matchin the provided filters.
// @Tags        Data Svc
// @Accept      json
// @Produce     json
// @Param       body      body     data.DeleteObjectRequest true "Delete request payload"
// @Success     200       {object} data.DeleteObjectResponse "Successful deletion of object"
// @Failure     400       {object} data.ErrorResponse "Invalid JSON"
// @Failure     401       {object} data.ErrorResponse "Unauthorized"
// @Failure     500       {object} data.ErrorResponse "Internal Server Error"
// @Security    BearerAuth
// @Router      /data-svc/objects/delete [post]
func (g *DataService) DeleteObjects(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, statusCode, err := g.options.PermissionChecker.HasPermission(
		r,
		data.PermissionObjectDelete,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.Authorized {
		endpoint.Unauthorized(w)
		return
	}

	req := &data.DeleteObjectRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		logger.Error("Error decoding request", slog.Any("error", err))
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	err = g.deleteObjects(req.Table, isAuthRsp.User.Id, req.Filters)
	if err != nil {
		logger.Error("Error deleting objects", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	_, err = w.Write([]byte(`{}`))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func (g *DataService) deleteObjects(
	tableName string,
	userId string,
	conditions []datastore.Filter,
) error {
	if len(conditions) == 0 {
		return errors.New("no conditions")
	}

	conditions = append(
		conditions,
		datastore.Equals(datastore.Field("table"), tableName),
	)
	conditions = append(conditions,
		datastore.Intersects(datastore.Field("deleters"), []any{userId}),
	)

	return g.store.Query(
		conditions...,
	).Delete()
}
