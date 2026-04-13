/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
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
// @Description Requires the `user-svc:organization:view` permission.
// @Description With `all=true`, platform admins see all organizations in the current app.
// @Description Otherwise users only see organizations they are members of.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.ListOrganizationsRequest false "List Organizations Request"
// @Success 200 {object} user.ListOrganizationsResponse "Organization listed successfully"
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/organizations [post]
func (s *UserService) ListOrganizations(
	w http.ResponseWriter,
	r *http.Request) {

	usr, hasPermission, claims, err := s.hasPermission(
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

	if r.Body != nil {
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
	}

	isAdmin := contains(claims.Roles, user.RoleAdmin)
	if req.All && !isAdmin {
		endpoint.Unauthorized(w)
		return
	}

	rsp, err := s.listOrganizations(
		claims.AppId,
		usr.Id,
		isAdmin,
		&req,
	)
	if err != nil {
		logger.Error(
			"Failed to list organizations",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	endpoint.WriteJSON(w, http.StatusOK, rsp)
}

func (s *UserService) listOrganizations(
	app string,
	userId string,
	isAdmin bool,
	request *user.ListOrganizationsRequest,
) (*user.ListOrganizationsResponse, error) {
	filters := []datastore.Filter{
		datastore.Equals(datastore.Field("appId"), app),
	}

	if !isAdmin || !request.All {
		membershipIs, err := s.membershipStore.Query(
			datastore.Equals(datastore.Field("appId"), app),
			datastore.Equals(datastore.Field("userId"), userId),
		).Find()
		if err != nil {
			return nil, err
		}

		if len(membershipIs) == 0 {
			return &user.ListOrganizationsResponse{
				Organizations: []user.Organization{},
			}, nil
		}

		organizationIds := []any{}
		for _, membershipI := range membershipIs {
			membership := membershipI.(*user.Membership)
			organizationIds = append(organizationIds, membership.OrganizationId)
		}

		filters = append(filters, datastore.IsInList(
			datastore.Field("id"),
			organizationIds...,
		))
	}

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

	q := s.organizationStore.Query(
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
