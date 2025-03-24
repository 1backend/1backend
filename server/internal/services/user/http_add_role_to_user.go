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

	sdk "github.com/1backend/1backend/sdk/go"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/gorilla/mux"
)

// @ID addRoleToUser
// @Summary Assign Role to User
// @Description Assign a role to a user. The caller can assign any roles it owns,
// @Description typically those prefixed with the caller’s identifier (e.g., `my-service:admin`).
// @Description One exception to this rule is dynamic organization roles: If the caller is an organization admin
// @Description (e.g., has a role like "user-svc:org:{%orgId}:admin"), they can also assign such roles.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Param roleId path string true "Role ID"
// @Param body body user.AddRoleToUserRequest true "Add Role to User Request"
// @Success 200 {object} user.AddRoleToUserResponse
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/user/{userId}/role/{roleId} [put]
func (s *UserService) AddRoleToUser(
	w http.ResponseWriter,
	r *http.Request) {

	authorizer := sdk.AuthorizerImpl{}
	claim, err := authorizer.ParseJWTFromRequest(s.publicKeyPem, r)
	if err != nil || claim == nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	// req := user.AddRoleToUserRequest{}
	// err = json.NewDecoder(r.Body).Decode(&req)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte(`Invalid JSON`))
	// 	return
	// }
	// defer r.Body.Close()

	userId := mux.Vars(r)["userId"]
	roleId := mux.Vars(r)["roleId"]

	if !sdk.OwnsRole(claim, roleId) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	err = s.addRoleToUser(userId, roleId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.AddRoleToUserResponse{})
	w.Write(bs)
}
