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
	"net/http"
	"time"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/datastore"
	user "github.com/1backend/1backend/server/internal/services/user/types"
)

// @ID saveInvites
// @Summary Save Invites
// @Description Invite a list of users by contact ID to acquire a role. Works on future or current users.
// @Description A user can only invite an other user to a role if the user owns that role.
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
// @Param body body user.SaveInvitesRequest true "Save Invites Request"
// @Success 200 {object} user.SaveInvitesResponse "Invites saved successfully"
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/invites [put]
func (s *UserService) SaveInvites(w http.ResponseWriter, r *http.Request) {

	usr, hasPermission, err := s.hasPermission(r, user.PermissionInviteEdit, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if !hasPermission {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	req := user.SaveInvitesRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	authr := auth.AuthorizerImpl{}
	claim, err := authr.ParseJWTFromRequest(s.publicKeyPem, r)
	if err != nil || claim == nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	for _, invite := range req.Invites {
		if !auth.OwnsRole(claim, invite.RoleId) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
	}

	invites, err := s.saveInvites(usr, &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.SaveInvitesResponse{
		Invites: invites,
	})
	w.Write(bs)
}

func (s *UserService) saveInvites(
	usr *user.User,
	req *user.SaveInvitesRequest,
) ([]user.Invite, error) {
	// @todo lock here

	if len(req.Invites) == 0 {
		return nil, errors.New("no invites provided")
	}
	now := time.Now()

	var (
		contactIds []any
		inviteIds  []any
	)
	for _, invite := range req.Invites {
		contactIds = append(contactIds, invite.ContactId)
		inviteIds = append(inviteIds, invite.Id)
	}

	contacts, err := s.contactsStore.Query(
		datastore.IsInList(
			datastore.Field("id"),
			contactIds...,
		)).
		Find()
	if err != nil {
		return nil, err
	}

	inviteIs, err := s.invitesStore.Query(
		datastore.IsInList(
			datastore.Field("id"),
			inviteIds...,
		)).
		Find()
	if err != nil {
		return nil, err
	}
	existingInvites := map[string]*user.Invite{}
	for _, inviteI := range inviteIs {
		existingInvites[inviteI.GetId()] = inviteI.(*user.Invite)
	}

	// Map contactIds -> userId
	existingContact := map[string]string{}
	for _, contact := range contacts {
		existingContact[contact.GetId()] = contact.(*user.Contact).UserId
	}

	invites := []user.Invite{}
	for _, invite := range req.Invites {
		// Already registered users get applied the role immediately
		if userId, ok := existingContact[invite.ContactId]; ok {
			err = s.assignRole(userId, invite.RoleId)
			if err != nil {
				return nil, err
			}
			continue
		}

		if invite.ContactId == "" || invite.RoleId == "" {
			return nil, errors.New("invite missing required fields")
		}

		if invite.Id == "" {
			invite.Id = sdk.Id("inv")
		}

		i := user.Invite{
			ContactId: invite.ContactId,
			RoleId:    invite.RoleId,
		}

		if existingInvite, ok := existingInvites[invite.Id]; ok {
			i.CreatedAt = existingInvite.CreatedAt
			i.UpdatedAt = now
			i.OwnerIds = append(existingInvite.OwnerIds, usr.Id)
		} else {
			i.CreatedAt = now
			i.UpdatedAt = now
			i.OwnerIds = []string{usr.Id}
		}

		invites = append(invites, i)
	}

	rows := []datastore.Row{}
	for _, invite := range invites {
		rows = append(rows, invite)
	}

	err = s.invitesStore.UpsertMany(rows)
	if err != nil {
		return nil, err
	}

	return invites, nil
}
