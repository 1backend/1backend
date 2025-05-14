/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package dynamicservice

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strings"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"

	data "github.com/1backend/1backend/server/internal/services/data/types"
)

// @ID createObject
// @Summary Create a Generic Object
// @Description Creates a new object with the provided details. Requires authorization and user authentication.
// @Tags Data Svc
// @Accept json
// @Produce json
// @Param body body data.CreateObjectRequest true "Create request payload"
// @Success 200 {object} data.CreateObjectResponse "Success"
// @Failure 400 {object} data.ErrorResponse "Invalid JSON"
// @Failure 401 {object} data.ErrorResponse "Unauthorized"
// @Failure 500 {object} data.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /data-svc/object [post]
func (g *DataService) Create(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, statusCode, err := g.permissionChecker.HasPermission(
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

	req := &data.CreateObjectRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		logger.Error(
			"Failed to decode request",
			slog.Any("error", err),
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	for i, v := range req.Object.Readers {
		if v == "_self" {
			req.Object.Readers[i] = isAuthRsp.User.Id
		}
	}
	for i, v := range req.Object.Writers {
		if v == "_self" {
			req.Object.Writers[i] = isAuthRsp.User.Id
		}
	}
	for i, v := range req.Object.Deleters {
		if v == "_self" {
			req.Object.Deleters[i] = isAuthRsp.User.Id
		}
	}

	err = g.createObject(req)
	if err != nil {
		logger.Error("Failed to create object", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	bs, _ := json.Marshal(data.CreateObjectResponse{
		Object: &data.Object{
			Id:    req.Object.Id,
			Table: req.Object.GetId(),
			Data:  req.Object.Data,
		},
	})
	_, err = w.Write(bs)
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func (g *DataService) createObject(
	request *data.CreateObjectRequest,
) error {
	if request.Object.Id == "" {
		request.Object.Id = sdk.Id(request.Object.Table)
	}
	if !strings.HasPrefix(request.Object.Id, request.Object.Table) {
		return errors.New("wrong prefix")
	}

	return g.store.Create(request.Object)
}
