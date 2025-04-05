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
	"time"

	"github.com/1backend/1backend/sdk/go/datastore"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	usertypes "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

// @ID saveUser
// @Summary Save User
// @Description Save user information based on the provided user ID.
// @Description It is intended for admins, because it uses the `user-svc:user:edit` permission which only admins have.
// @Description For a user to edit its own profile, see saveSelf.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Param body body user.SaveProfileRequest true "Save Profile Request"
// @Success 200 {object} user.SaveProfileResponse
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/user/{userId} [put]
func (s *UserService) SaveUser(w http.ResponseWriter, r *http.Request) {

	_, isAuthorized, err := s.hasPermission(r, user.PermissionUserEdit, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if !isAuthorized {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	vars := mux.Vars(r)
	userId := vars["userId"]

	req := user.SaveProfileRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	err = s.saveUser(userId, &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.SaveProfileResponse{})
	w.Write(bs)
}

func (s *UserService) saveUser(
	userId string,
	request *user.SaveProfileRequest,
) error {
	query := s.usersStore.Query(
		datastore.Equals(datastore.Field("id"), userId),
	)

	userI, found, err := query.FindOne()
	if err != nil {
		return err
	}

	if !found {
		return errors.New("user not found")
	}
	user := userI.(*usertypes.User)

	if request.Name != "" {
		user.Name = request.Name
	}

	if request.ThumbnailFileId != "" {
		user.ThumbnailFileId = request.ThumbnailFileId
	}

	user.UpdatedAt = time.Now()

	query.Update(user)

	return nil
}
