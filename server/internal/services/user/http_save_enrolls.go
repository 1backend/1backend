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
	"errors"
	"log/slog"
	"net/http"
	"time"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
)

// @ID saveEnrolls
// @Summary Save Enrolls
// @Description Enroll a list of users by contact or user Id to acquire a role.
// @Description Works on future or current users.
// @Description
// @Description A user can only enroll an other user to a role if the user owns that role.
// @Description
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
// @Param body body user.SaveEnrollsRequest true "Save Enrolls Request"
// @Success 200 {object} user.SaveEnrollsResponse "Enrolls saved successfully"
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/enrolls [put]
func (s *UserService) SaveEnrolls(w http.ResponseWriter, r *http.Request) {

	usr, hasPermission, err := s.hasPermission(r, user.PermissionEnrollEdit)
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

	req := user.SaveEnrollsRequest{}
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

	isAdmin, err := s.authorizer.IsAdminFromRequest(s.publicKeyPem, r)
	if err != nil {
		logger.Error(
			"Failed to check if user is admin",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	claim, err := s.parseJWTFromRequest(r)
	if err != nil || claim == nil {
		endpoint.Unauthorized(w)
		return
	}

	if !isAdmin {
		for _, enroll := range req.Enrolls {
			if !auth.OwnsRole(claim, enroll.Role) {
				endpoint.Unauthorized(w)
				return
			}
		}
	}

	enrolls, err := s.saveEnrolls(usr.Id, &req)
	if err != nil {
		logger.Error(
			"Failed to save enrolls",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	bs, _ := json.Marshal(user.SaveEnrollsResponse{
		Enrolls: enrolls,
	})
	_, err = w.Write(bs)
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func (s *UserService) saveEnrolls(
	callerUserId string,
	req *user.SaveEnrollsRequest,
) ([]user.Enroll, error) {
	// @todo lock here

	if len(req.Enrolls) == 0 {
		return nil, errors.New("no enrolls provided")
	}
	now := time.Now()

	var (
		contactIds    []any
		callerUserIds []any
		enrollIds     []any
	)
	for _, enroll := range req.Enrolls {
		contactIds = append(contactIds, enroll.ContactId)
		enrollIds = append(enrollIds, enroll.Id)
		callerUserIds = append(callerUserIds, enroll.UserId)
	}

	var (
		err      error
		contacts []datastore.Row
		users    []datastore.Row
		enrollIs []datastore.Row
	)

	if len(contactIds) > 0 {
		contacts, err = s.contactsStore.Query(
			datastore.IsInList(
				datastore.Field("id"),
				contactIds...,
			)).
			Find()
		if err != nil {
			return nil, err
		}
	}

	if len(callerUserIds) > 0 {
		users, err = s.usersStore.Query(
			datastore.IsInList(
				datastore.Field("id"),
				callerUserIds...,
			)).
			Find()
		if err != nil {
			return nil, err
		}
	}

	if len(enrollIds) > 0 {
		enrollIs, err = s.enrollsStore.Query(
			datastore.IsInList(
				datastore.Field("id"),
				enrollIds...,
			)).
			Find()
		if err != nil {
			return nil, err
		}
	}

	existingEnrolls := map[string]*user.Enroll{}
	for _, enrollI := range enrollIs {
		existingEnrolls[enrollI.GetId()] = enrollI.(*user.Enroll)
	}

	// Map contactIds -> callerUserId
	existingContact := map[string]string{}
	for _, contact := range contacts {
		existingContact[contact.GetId()] = contact.(*user.Contact).UserId
	}

	existingUser := map[string]bool{}
	for _, usr := range users {
		existingUser[usr.GetId()] = true
	}

	enrolls := []user.Enroll{}
	for _, enroll := range req.Enrolls {
		if enroll.Id == "" {
			enroll.Id = sdk.Id("enr")
		}

		// Already registered users get applied the role immediately
		if callerUserId, ok := existingContact[enroll.ContactId]; ok {
			err = s.assignRole(callerUserId, enroll.Role)
			if err != nil {
				return nil, err
			}
			continue
		}

		if _, ok := existingUser[enroll.UserId]; ok {
			err = s.assignRole(enroll.UserId, enroll.Role)
			if err != nil {
				return nil, err
			}
			continue
		}

		if (enroll.ContactId == "" && enroll.UserId == "") || enroll.Role == "" {
			return nil, errors.New("enroll missing required fields")
		}

		i := user.Enroll{
			Id:        enroll.Id,
			ContactId: enroll.ContactId,
			Role:      enroll.Role,
			CreatedBy: callerUserId,
		}

		if existingEnroll, ok := existingEnrolls[enroll.Id]; ok {
			i.CreatedAt = existingEnroll.CreatedAt
			i.UpdatedAt = now
		} else {
			i.CreatedAt = now
			i.UpdatedAt = now
		}

		enrolls = append(enrolls, i)
	}

	rows := []datastore.Row{}
	for _, enroll := range enrolls {
		rows = append(rows, enroll)
	}

	err = s.enrollsStore.UpsertMany(rows)
	if err != nil {
		return nil, err
	}

	return enrolls, nil
}
