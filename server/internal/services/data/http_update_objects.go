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
	"net/http"

	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/datastore"
	data "github.com/openorch/openorch/server/internal/services/data/types"
	dynamictypes "github.com/openorch/openorch/server/internal/services/data/types"
)

// @ID updateObjects
// @Summary Update Objects
// @Description Update fields of objects that match the given filters using the provided object.
// @Description Any fields not included in the incoming object will remain unchanged.
// @Tags Data Svc
// @Accept json
// @Produce json
// @Param body body data.UpdateObjectsRequest true "Update request payload"
// @Success 200 {object} data.UpdateObjectsResponse "Successful update of objects"
// @Failure 400 {object} data.ErrorResponse "Invalid JSON"
// @Failure 401 {object} data.ErrorResponse "Unauthorized"
// @Failure 500 {object} data.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /data-svc/objects/update [post]
func (g *DataService) UpdateObjects(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, _, err := g.clientFactory.Client(sdk.WithTokenFromRequest(r)).
		UserSvcAPI.IsAuthorized(r.Context(), *data.PermissionObjectEdit.Id).
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

	req := &data.UpdateObjectsRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	err = g.updateObjects(req.Table, *isAuthRsp.User.Id, req.Filters, req.Object)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(`{}`))
}

func (g *DataService) updateObjects(
	tableName string,
	userId string,
	conditions []datastore.Filter,
	object *dynamictypes.Object,
) error {
	if len(conditions) == 0 {
		return errors.New("no conditions")
	}

	conditions = append(
		conditions,
		datastore.Equals(datastore.Field("table"), tableName),
	)
	conditions = append(conditions, datastore.Equals(
		datastore.Field("userId"),
		userId,
	))

	return g.store.Query(
		conditions...,
	).Update(object)
}
