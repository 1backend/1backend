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
	"github.com/1backend/1backend/sdk/go/datastore"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/samber/lo"
)

// @ID saveInvites
// @Summary Save Invites
// @Description Save a list of user invites to the database.
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

	_, err := s.isAuthorized(r, user.PermissionInviteEdit.Id, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
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

	auth := sdk.AuthorizerImpl{}
	claim, err := auth.ParseJWTFromRequest(s.publicKeyPem, r)
	if err != nil || claim == nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	for _, invite := range req.Invites {
		if !sdk.OwnsRole(claim, invite.RoleId) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
	}

	invites, err := s.saveInvites(&req)
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
	req *user.SaveInvitesRequest,
) ([]user.Invite, error) {

	if len(req.Invites) == 0 {
		return nil, errors.New("no invites provided")
	}
	now := time.Now()

	contactIds := lo.Map(req.Invites, func(inv user.NewInvite, _ int) any {
		return inv.ContactId
	})

	contacts, err := s.contactsStore.Query(
		datastore.IsInList(
			datastore.Field("id"),
			contactIds...,
		)).
		Find()
	if err != nil {
		return nil, err
	}

	// Map contactIds -> userId
	existingContact := map[string]string{}
	for _, contact := range contacts {
		existingContact[contact.GetId()] = contact.(*user.Contact).UserId
	}

	invites := []user.Invite{}
	for _, invite := range req.Invites {
		if userId, ok := existingContact[invite.ContactId]; ok {
			err = s.addRoleToUser(userId, invite.RoleId)
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

		invites = append(invites, user.Invite{
			CreatedAt: now,
			ContactId: invite.ContactId,
			RoleId:    invite.RoleId,
		})
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
