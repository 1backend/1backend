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
	"net/url"

	"github.com/1backend/1backend/sdk/go/auth"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/gorilla/mux"
)

// @ID assignRole
// @Summary Assign Role
// @Description Assigns a role to a user. The caller can only assign roles they own.
// @Description A user "owns" a role in the following cases:
// @Description - A static role where the role ID is prefixed with the caller's slug.
// @Description - Any dynamic or static role where the caller is an admin.
// @Description
// @Description Examples:
// @Description - A user with the slug "joe-doe" owns roles like "joe-doe:any-custom-role".
// @Description - A user with any slug who has the role "my-service:admin" owns "my-service:user".
// @Description - A user with any slug who has the role "user-svc:org:{%orgId}:admin" owns "user-svc:org:{%orgId}:user".
// @Tags User Svc
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Param role path string true "Role ID"
// @Param body body user.AssignRoleRequest false "Assign Role Request"
// @Success 200 {object} user.AssignRoleResponse
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Role not found"
// @Security BearerAuth
// @Router /user-svc/user/{userId}/role/{role} [put]
func (s *UserService) AssignRole(
	w http.ResponseWriter,
	r *http.Request) {

	authorizer := auth.AuthorizerImpl{}
	claim, err := authorizer.ParseJWTFromRequest(s.publicKeyPem, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if claim == nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	// req := user.AssignRoleRequest{}
	// err = json.NewDecoder(r.Body).Decode(&req)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte(`Invalid JSON`))
	// 	return
	// }
	// defer r.Body.Close()

	userId, err := url.PathUnescape(mux.Vars(r)["userId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	role, err := url.PathUnescape(mux.Vars(r)["role"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}

	if !auth.OwnsRole(claim, role) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	err = s.assignRole(userId, role)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.AssignRoleResponse{})
	w.Write(bs)
}
