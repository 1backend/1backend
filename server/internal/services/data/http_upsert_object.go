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

	"github.com/gorilla/mux"

	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/datastore"
	data "github.com/openorch/openorch/server/internal/services/data/types"
)

// @ID upsertObject
// @Summary Upsert a Generic Object
// @Description Creates a new dynamic object or updates an existing one based on the provided data. Requires authorization and user authentication.
// @Tags Data Svc
// @Accept json
// @Produce json
// @Param objectId path string true  "Object ID"
// @Param body body data.UpsertObjectRequest true "Upsert request payload"
// @Success 200 {object} data.UpsertObjectResponse "Successful creation or update of object"
// @Failure 400 {object} data.ErrorResponse "Invalid JSON"
// @Failure 401 {object} data.ErrorResponse "Unauthorized"
// @Failure 500 {object} data.ErrorResponse "Internal Server Error"
// @Security    BearerAuth
// @Router /data-svc/object/{objectId} [put]
func (g *DataService) Upsert(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, _, err := g.clientFactory.Client(sdk.WithTokenFromRequest(r)).
		UserSvcAPI.IsAuthorized(r.Context(), *data.PermissionObjectCreate.Id).
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

	req := &data.UpsertObjectRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	for i, v := range req.Object.Readers {
		if v == "_self" {
			req.Object.Readers[i] = *isAuthRsp.User.Id
		}
	}
	for i, v := range req.Object.Writers {
		if v == "_self" {
			req.Object.Writers[i] = *isAuthRsp.User.Id
		}
	}
	for i, v := range req.Object.Deleters {
		if v == "_self" {
			req.Object.Deleters[i] = *isAuthRsp.User.Id
		}
	}

	claims, err := g.authorizer.ParseJWTFromRequest(g.publicKey, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	identifiers := append(claims.RoleIds, *isAuthRsp.User.Id)

	objectId := mux.Vars(r)
	req.Object.Id = objectId["objectId"]

	err = g.upsertObject(identifiers, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(`{}`))
}

func (g *DataService) upsertObject(
	writers []string,
	request *data.UpsertObjectRequest,
) error {
	vI, found, err := g.store.Query(
		datastore.Id(request.Object.Id),
	).FindOne()
	if err != nil {
		return err
	}

	if found {
		v := vI.(*data.Object)

		if !intersects(writers, v.Writers) {
			return errors.New("unauthorized")
		}

		if request.Object.Readers == nil {
			request.Object.Readers = v.Readers
		}
		if request.Object.Writers == nil {
			request.Object.Writers = v.Writers
		}
		if request.Object.Deleters == nil {
			request.Object.Deleters = v.Deleters
		}
	}

	return g.store.Upsert(request.Object)
}
