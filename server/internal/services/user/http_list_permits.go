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

// @ID listPermits
// @Summary List Permits
// @Description
// @Description List permits. Requires the `user-svc:permit:view` permission, which only admins have by default.
// @Description &todo Users should be able to list permits referring to them.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.ListPermitsRequest true "List Permits Request"
// @Success 200 {object} user.ListPermitsResponse
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/permits [post]
func (s *UserService) ListPermits(
	w http.ResponseWriter,
	r *http.Request) {

	_, has, err := s.hasPermission(r, user.PermissionPermitView)
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

	req := &user.ListPermitsRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	permits, err := s.listPermits(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.ListPermitsResponse{
		Permits: permits,
	})
	w.Write(bs)
}

func (us *UserService) listPermits(req *user.ListPermitsRequest) ([]*user.Permit, error) {
	filters := []datastore.Filter{}
	if req.Permission != "" {
		filters = append(filters, datastore.Equals([]string{"permission"}, req.Permission))
	}
	if req.Slug != "" {
		filters = append(filters, datastore.Equals([]string{"slug"}, req.Slug))
	}

	permitIs, err := us.permitsStore.Query(filters...).Find()
	if err != nil {
		return nil, errors.Wrap(err, "error querying permits")
	}

	permits := []*user.Permit{}
	for _, permitI := range permitIs {
		permits = append(permits, permitI.(*user.Permit))
	}

	return permits, nil
}
