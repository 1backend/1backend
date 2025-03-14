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
	"context"
	"encoding/json"
	"net/http"

	sdk "github.com/openorch/openorch/sdk/go"
	user "github.com/openorch/openorch/server/internal/services/user/types"
	"github.com/pkg/errors"
)

// @ID saveGrants
// @Summary Save Grants
// @Description Save grants.
// @Description
// @Description Grants define which slugs are assigned specific permissions, overriding the default configuration.
// @Description
// @Description Requires the `user-svc:grant:create` permission.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.SaveGrantsRequest true "Save Grants Request"
// @Success 200 {object} user.SaveGrantsResponse "Grant saved successfully"
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/grants [put]
func (s *UserService) SaveGrants(w http.ResponseWriter, r *http.Request) {

	_, err := s.isAuthorized(r, user.PermissionRoleCreate.Id, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	req := user.SaveGrantsRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	err = s.saveGrants(
		r.Context(),
		&req,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.SaveGrantsResponse{})
	w.Write(bs)
}

func (cs *UserService) saveGrants(
	ctx context.Context,
	req *user.SaveGrantsRequest,
) error {
	for _, grant := range req.Grants {
		if grant.Id == "" {
			grant.Id = sdk.Id("grn")
		}

		err := cs.grantsStore.Upsert(grant)
		if err != nil {
			return errors.Wrap(err, "error saving grant")
		}
	}

	return nil
}
