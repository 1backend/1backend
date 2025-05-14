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

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
)

// @ID listOrganizations
// @Summary List Organizations
// @Description Requires the `user-svc:organization:view` permission, that only admins have by default.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.ListOrganizationsRequest true "List Organizations Request"
// @Success 200 {object} user.ListOrganizationsResponse "Organization listed successfully"
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/organizations [post]
func (s *UserService) ListOrganizations(
	w http.ResponseWriter,
	r *http.Request) {

	_, hasPermission, err := s.hasPermission(
		r,
		user.PermissionOrganizationView,
	)
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

	req := user.ListOrganizationsRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error(
			"Failed to decode request",
			slog.Any("error", err),
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	rsp, err := s.listOrganizations(&req)
	if err != nil {
		logger.Error(
			"Failed to list organizations",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	bs, _ := json.Marshal(rsp)
	w.Write(bs)
}

func (s *UserService) listOrganizations(
	request *user.ListOrganizationsRequest,
) (*user.ListOrganizationsResponse, error) {
	filters := []datastore.Filter{}

	if request.Ids != nil {
		ids := []any{}
		for _, id := range request.Ids {
			ids = append(ids, id)
		}

		filters = append(filters, datastore.IsInList(
			[]string{"id"},
			ids...,
		))
	}

	q := s.organizationsStore.Query(
		filters...,
	).OrderBy(
		datastore.OrderByField("createdAt", false),
	)

	if !request.AfterTime.IsZero() {
		q = q.After(request.AfterTime)
	}

	if request.Limit != 0 {
		q = q.Limit(int64(request.Limit))
	} else {
		q = q.Limit(20)
	}

	organizationIs, err := q.Find()
	if err != nil {
		return nil, err
	}

	organizations := []user.Organization{}
	for _, v := range organizationIs {
		organizations = append(organizations, *v.(*user.Organization))
	}

	return &user.ListOrganizationsResponse{
		Organizations: organizations,
	}, nil

}
