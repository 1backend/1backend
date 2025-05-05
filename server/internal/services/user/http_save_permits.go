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
	"strings"
	"time"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/datastore"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/pkg/errors"
)

// @ID savePermits
// @Summary Save Permits
// @Description Save permits. // @Description Permits give access to users with certain slugs and roles to permissions.
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
	usr, err := s.getUserFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	req := user.SavePermitsRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	isAdmin, err := auth.AuthorizerImpl{}.IsAdminFromRequest(s.publicKeyPem, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if !isAdmin {
		for _, permit := range req.Permits {
			if !strings.HasPrefix(permit.Permission, usr.Slug) {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
				return
			}
		}
	}

	err = s.savePermits(
		r.Context(),
		&req,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.SavePermitsResponse{})
	w.Write(bs)
}

func (cs *UserService) savePermits(
	ctx context.Context,
	req *user.SavePermitsRequest,
) error {
	permissions := []any{}
	for _, permit := range req.Permits {
		permissions = append(permissions, permit.Permission)
	}

	permitIs, err := cs.permitsStore.Query(
		datastore.IsInList([]string{"permission"}, permissions...),
	).Find()
	if err != nil {
		return errors.Wrap(err, "failed to list permits by permission")
	}

	permitsByPermission := map[string][]*user.Permit{}
	for _, permitI := range permitIs {
		g := permitI.(*user.Permit)
		permitsByPermission[g.Permission] = append(permitsByPermission[g.Permission], g)
	}

	permits := []datastore.Row{}

	now := time.Now()
	for _, permit := range req.Permits {
		nu := false
		if permit.Id == "" {
			permit.Id = sdk.Id("pmt")
			nu = true
		}

		existingPermits, ok := permitsByPermission[permit.Permission]
		isDuplicate := false
		if ok {
			for _, existingPermit := range existingPermits {
				if equalUnordered(existingPermit.Roles, permit.Roles) &&
					equalUnordered(existingPermit.Slugs, permit.Slugs) {
					isDuplicate = true
					break
				}
			}
		}
		if isDuplicate {
			continue
		}

		permit := &user.Permit{
			Id:         permit.Id,
			Permission: permit.Permission,
			Slugs:      permit.Slugs,
			Roles:      permit.Roles,
			UpdatedAt:  now,
		}
		if nu {
			permit.CreatedAt = now
		}

		permits = append(permits, permit)
	}

	if len(permits) == 0 {
		return nil
	}

	err = cs.permitsStore.UpsertMany(permits)
	if err != nil {
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
