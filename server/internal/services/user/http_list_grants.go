/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package userservice

import (
	"encoding/json"
	"net/http"

	"github.com/1backend/1backend/sdk/go/datastore"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/pkg/errors"
)

// @ID listGrants
// @Summary List Grants
// @Description
// @Description Grants give access to users with certain slugs and roles to permissions.
// @Description Users can list grants for permissions they have access to
// @Description but they will only see grants the grant refers to their slug or one of their roles.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.ListGrantsRequest true "List Grants Request"
// @Success 200 {object} user.ListGrantsResponse
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/grants [post]
func (s *UserService) ListGrants(
	w http.ResponseWriter,
	r *http.Request) {

	req := &user.ListGrantsRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	if req.Permission != "" {
		_, has, err := s.hasPermission(r, req.Permission, nil, nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		if !has {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`Unauthorized`))
			return
		}
	}

	grants, err := s.listGrants(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.ListGrantsResponse{
		Grants: grants,
	})
	w.Write(bs)
}

func (us *UserService) listGrants(req *user.ListGrantsRequest) ([]*user.Grant, error) {
	filters := []datastore.Filter{}
	if req.Permission != "" {
		filters = append(filters, datastore.Equals([]string{"permission"}, req.Permission))
	}
	if req.Slug != "" {
		filters = append(filters, datastore.Equals([]string{"slug"}, req.Slug))
	}

	grantIs, err := us.grantsStore.Query(filters...).Find()
	if err != nil {
		return nil, errors.Wrap(err, "error querying grants")
	}

	grants := []*user.Grant{}
	for _, grantI := range grantIs {
		grants = append(grants, grantI.(*user.Grant))
	}

	return grants, nil
}
