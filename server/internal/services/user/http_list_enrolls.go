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
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

// @ID listEnrolls
// @Summary List Enrolls
// @Description List enrolls. Role, user ID or contact ID must be specified.
// @Description Caller can only list enrolls of roles they own.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.ListEnrollsRequest true "List Enrolls Request"
// @Success 200 {object} user.ListEnrollsResponse "Enrolls listed successfully"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/enrolls [post]
func (s *UserService) ListEnrolls(w http.ResponseWriter, r *http.Request) {

	_, hasPermission, err := s.hasPermission(r, user.PermissionEnrollView)
	if err != nil {
		logger.Error(
			"Failed to check permission",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}
	if !hasPermission {
		endpoint.Unauthorized(w)
		return
	}

	authr := auth.AuthorizerImpl{}
	claim, err := authr.ParseJWTFromRequest(s.publicKeyPem, r)
	if err != nil || claim == nil {
		endpoint.Unauthorized(w)
		return
	}

	req := &user.ListEnrollsRequest{}
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

	enrolls, err := s.listEnrolls(claim, req)
	if err != nil {
		if err != nil {
			logger.Error(
				"Failed to list enrolls",
				slog.Any("error", err),
			)
			endpoint.InternalServerError(w)
			return
		}
		if !hasPermission {
			endpoint.Unauthorized(w)
			return
		}
		return
	}

	bs, _ := json.Marshal(user.ListEnrollsResponse{
		Enrolls: enrolls,
	})
	w.Write(bs)
}

func (s *UserService) listEnrolls(
	claims *auth.Claims,
	req *user.ListEnrollsRequest,
) ([]user.Enroll, error) {
	if req.Role == "" && req.ContactId == "" && req.UserId == "" {
		return nil, errors.New("role, contact id or user id is required")
	}

	isAdmin := lo.Contains(claims.Roles, user.RoleAdmin)

	filters := []datastore.Filter{}

	if req.Role != "" {
		if !isAdmin {
			if !auth.OwnsRole(claims, req.Role) {
				return nil, errors.New("cannot list enrolls for unowned role")
			}
		}
		filters = append(filters, datastore.Equals([]string{"role"}, req.Role))
	}
	if req.ContactId != "" {
		if !isAdmin {
			// @todo let users list for their own contacts
			return nil, errors.New("only admins can list based on contact id")
		}
		filters = append(filters, datastore.Equals([]string{"contactId"}, req.ContactId))
	}
	if req.UserId != "" {
		if !isAdmin && req.UserId != claims.UserId {
			return nil, errors.New("cannot list for other user ids")
		}
		filters = append(filters, datastore.Equals([]string{"userId"}, req.UserId))
	}

	enrollIs, err := s.enrollsStore.Query(filters...).Find()
	if err != nil {
		return nil, errors.Wrap(err, "error querying enrolls")
	}

	enrolls := []user.Enroll{}
	for _, enrollI := range enrollIs {
		e := *enrollI.(*user.Enroll)

		if isAdmin || auth.OwnsRole(claims, e.Role) {
			enrolls = append(enrolls, e)
		}
	}

	return enrolls, nil
}
