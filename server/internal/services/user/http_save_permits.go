/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/pkg/errors"
)

// @ID savePermits
// @Summary Save Permits
// @Description Save permits.
// @Description Permits give access to users with certain slugs and roles to permissions.
// @Description Non-admin callers may only save permissions they own, such as `their-slug:...` or organization-scoped namespaces they administer.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.SavePermitsRequest true "Save Permits Request"
// @Success 200 {object} user.SavePermitsResponse "Permit saved successfully"
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/permits [put]
func (s *UserService) SavePermits(w http.ResponseWriter, r *http.Request) {
	if _, hasToken := s.options.Authorizer.TokenFromRequest(r); !hasToken {
		endpoint.Unauthorized(w)
		return
	}

	_, claims, err := s.getUserFromRequest(r)
	if err != nil {
		if isUnauthorizedRequestError(err) {
			endpoint.Unauthorized(w)
			return
		}

		logger.Error(
			"Failed to get user from request",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	req := user.SavePermitsRequest{}
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

	for _, permit := range req.Permits {
		permissions := normalizePermitInputPermissions(permit)
		if len(permissions) == 0 {
			endpoint.WriteString(w, http.StatusBadRequest, "Permit must include permission or permissions")
			return
		}

		for _, permission := range permissions {
			if !auth.OwnsPermission(claims, permission) {
				endpoint.Unauthorized(w)
				return
			}
		}
	}

	err = s.savePermits(
		claims.AppId,
		r.Context(),
		&req,
	)
	if err != nil {
		logger.Error(
			"Failed to save permits",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	bs, _ := json.Marshal(user.SavePermitsResponse{})
	_, err = w.Write(bs)
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func (cs *UserService) savePermits(
	claimAppId string,
	ctx context.Context,
	req *user.SavePermitsRequest,
) error {
	permissions := []any{}
	for _, permit := range req.Permits {
		for _, permission := range normalizePermitInputPermissions(permit) {
			permissions = append(permissions, permission)
		}
	}

	existingPermits := []*user.Permit{}
	if len(permissions) > 0 {
		permitIs, err := cs.permitStore.Query(
			datastore.Or(
				datastore.IsInList([]string{"permission"}, permissions...),
				datastore.Intersects([]string{"permissions"}, permissions),
			),
		).Find()
		if err != nil {
			return errors.Wrap(err, "failed to list permits by permission")
		}

		for _, permitI := range permitIs {
			existingPermits = append(existingPermits, permitI.(*user.Permit))
		}
	}

	permits := []datastore.Row{}
	seenPermitKeys := map[string]struct{}{}

	now := time.Now()
	for _, permit := range req.Permits {
		normalizedPermissions := normalizePermitInputPermissions(permit)
		if len(normalizedPermissions) == 0 {
			continue
		}

		nu := false
		if permit.Id == "" {
			permit.Id = sdk.Id("perm")
			nu = true
		}

		permitAppId := ""
		if permit.AppHost != "" && permit.AppHost != "*" {
			app, found, err := cs.appByHost(permit.AppHost)
			if err != nil {
				return errors.Wrap(err, "failed to query app by host")
			}
			if !found {
				return errors.Errorf("app with host %s not found", permit.AppHost)
			}
			permitAppId = app.Id
		} else if permit.AppHost == "*" {
			permitAppId = "*"
		} else {
			permitAppId = claimAppId
		}

		identityKey := permitIdentityKey(
			permitAppId,
			permit.Roles,
			permit.Slugs,
			normalizedPermissions,
		)
		if _, ok := seenPermitKeys[identityKey]; ok {
			continue
		}

		isDuplicate := false
		for _, existingPermit := range existingPermits {
			if existingPermit.AppId != permitAppId {
				continue
			}
			if !equalUnordered(existingPermit.Roles, permit.Roles) {
				continue
			}
			if !equalUnordered(existingPermit.Slugs, permit.Slugs) {
				continue
			}
			if !equalUnordered(permitPermissions(existingPermit), normalizedPermissions) {
				continue
			}

			isDuplicate = true
			break
		}
		if isDuplicate {
			seenPermitKeys[identityKey] = struct{}{}
			continue
		}

		internalId, err := sdk.InternalId(permitAppId, permit.Id)
		if err != nil {
			return errors.Wrap(err, "failed to create internal id")
		}

		permissionField, permissionsField := canonicalPermissionFields(normalizedPermissions)

		permit := &user.Permit{
			InternalId:  internalId,
			Id:          permit.Id,
			AppId:       permitAppId,
			Permission:  permissionField,
			Permissions: permissionsField,
			Slugs:       permit.Slugs,
			Roles:       permit.Roles,
			UpdatedAt:   now,
		}
		if nu {
			permit.CreatedAt = now
		}

		existingPermits = append(existingPermits, permit)
		seenPermitKeys[identityKey] = struct{}{}
		permits = append(permits, permit)
	}

	if len(permits) == 0 {
		return nil
	}

	if err := cs.permitStore.UpsertMany(permits); err != nil {
		return errors.Wrap(err, "error saving permits")
	}

	return nil
}

// equalUnordered checks if two slices contain the same elements regardless of order.
// Assumes elements are comparable (e.g., int, string, etc.).
func equalUnordered[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	counts := make(map[T]int)

	for _, item := range a {
		counts[item]++
	}
	for _, item := range b {
		counts[item]--
		if counts[item] < 0 {
			return false
		}
	}

	return true
}
