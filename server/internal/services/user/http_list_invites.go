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

	"github.com/1backend/1backend/sdk/go/datastore"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/pkg/errors"
)

// @ID listInvites
// @Summary List Invites
// @Description List user invites stored in the database.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.ListInvitesRequest true "List Invites Request"
// @Success 200 {object} user.ListInvitesResponse "Invites listed successfully"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/invites [post]
func (s *UserService) ListInvites(w http.ResponseWriter, r *http.Request) {

	usr, isAuthorized, err := s.isAuthorized(r, user.PermissionInviteView.Id, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if !isAuthorized {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	req := &user.ListInvitesRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	invites, err := s.listInvites(usr.Id, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(user.ListInvitesResponse{
		Invites: invites,
	})
	w.Write(bs)
}

func (s *UserService) listInvites(
	callerId string,
	req *user.ListInvitesRequest,
) ([]user.Invite, error) {
	filters := []datastore.Filter{
		datastore.Equals([]string{"ownerIds"}, callerId),
	}
	if req.ContactId != "" {
		filters = append(filters, datastore.Equals([]string{"contactId"}, req.ContactId))
	}
	if req.RoleId != "" {
		filters = append(filters, datastore.Equals([]string{"roleId"}, req.RoleId))
	}

	inviteIs, err := s.invitesStore.Query(filters...).Find()
	if err != nil {
		return nil, errors.Wrap(err, "error querying invites")
	}

	invites := []user.Invite{}
	for _, inviteI := range inviteIs {
		i := *inviteI.(*user.Invite)
		i.OwnerIds = nil
		invites = append(invites, i)
	}

	return invites, nil
}
