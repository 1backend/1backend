/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 */
package userservice

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

var (
	cannotDeleteUnownedEnroll = errors.New("cannot delete enroll for unowned role")
	enrollNotFound            = errors.New("enroll not found")
)

// @ID deleteEnrolls
// @Summary Delete Enrolls
// @Description Delete enrolls by ID.
// @Description Requires the `user-svc:enroll:edit` permission, which by default all users have.
// @Description Caller can only delete enrolls of roles they own (unless they are an admin).
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.DeleteEnrollsRequest true "Delete Enrolls Request"
// @Success 200 {object} user.DeleteEnrollsResponse "Enrolls deleted successfully"
// @Failure 400 {object} user.ErrorResponse "Invalid JSON or missing IDs"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 404 {object} user.ErrorResponse "Enroll not found"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/enrolls [delete]
func (s *UserService) DeleteEnrolls(w http.ResponseWriter, r *http.Request) {
	_, hasPermission, claims, err := s.hasPermission(r, user.PermissionEnrollEdit)
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

	req := &user.DeleteEnrollsRequest{}
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

	req.Ids = uniqueEnrollIds(req.Ids)
	if len(req.Ids) == 0 {
		endpoint.WriteString(w, http.StatusBadRequest, "IDs are required")
		return
	}

	roles, err := s.getRolesByUserId(claims.AppId, claims.UserId)
	if err != nil {
		logger.Error("Failed to list effective roles", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	effectiveClaims := *claims
	effectiveClaims.Roles = roles

	err = s.deleteEnrolls(&effectiveClaims, req)
	switch {
	case errors.Is(err, cannotDeleteUnownedEnroll):
		endpoint.Unauthorized(w)
		return
	case errors.Is(err, enrollNotFound):
		endpoint.WriteString(w, http.StatusNotFound, "Enroll not found")
		return
	case err != nil:
		logger.Error(
			"Failed to delete enrolls",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	endpoint.WriteJSON(w, http.StatusOK, user.DeleteEnrollsResponse{})
}

func (s *UserService) deleteEnrolls(
	claims *auth.Claims,
	req *user.DeleteEnrollsRequest,
) error {
	ids := uniqueEnrollIds(req.Ids)
	if len(ids) == 0 {
		return errors.New("no enroll IDs provided")
	}

	idAnys := make([]any, 0, len(ids))
	for _, id := range ids {
		idAnys = append(idAnys, id)
	}

	enrollIs, err := s.enrollStore.Query(
		datastore.IsInList(datastore.Field("id"), idAnys...),
	).Find()
	if err != nil {
		return errors.Wrap(err, "failed to query enrolls")
	}
	if len(enrollIs) != len(ids) {
		return enrollNotFound
	}

	isAdmin := lo.Contains(claims.Roles, user.RoleAdmin)
	userIdsByAppIdToRefresh := map[string]map[string]struct{}{}
	for _, enrollI := range enrollIs {
		enroll := enrollI.(*user.Enroll)
		if !isAdmin && !auth.OwnsRole(claims, enroll.Role) {
			return cannotDeleteUnownedEnroll
		}
		if enroll.UserId != "" {
			if userIdsByAppIdToRefresh[enroll.AppId] == nil {
				userIdsByAppIdToRefresh[enroll.AppId] = map[string]struct{}{}
			}
			userIdsByAppIdToRefresh[enroll.AppId][enroll.UserId] = struct{}{}
		}
	}

	if err := s.enrollStore.Query(
		datastore.IsInList(datastore.Field("id"), idAnys...),
	).Delete(); err != nil {
		return errors.Wrap(err, "failed to delete enrolls")
	}

	for appId, userIds := range userIdsByAppIdToRefresh {
		for userId := range userIds {
			if err := s.inactivateTokens(appId, userId); err != nil {
				return err
			}
		}
	}

	return nil
}

func uniqueEnrollIds(ids []string) []string {
	seen := map[string]struct{}{}
	unique := []string{}
	for _, id := range ids {
		id = strings.TrimSpace(id)
		if id == "" {
			continue
		}
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}
		unique = append(unique, id)
	}
	return unique
}
