package userservice

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

var ErrPendingMembershipNotFound = errors.New("pending organization membership not found")

// @ID acceptMembership
// @Summary Accept Membership
// @Description Accepts the caller user's pending invite for an organization.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param organizationId path string true "Organization ID"
// @Param body body user.AcceptMembershipRequest false "Accept Membership Request"
// @Success 200 {object} user.AcceptMembershipResponse
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/organization/{organizationId}/membership/accept [post]
func (s *UserService) AcceptMembership(
	w http.ResponseWriter,
	r *http.Request,
) {
	organizationId := mux.Vars(r)["organizationId"]

	usr, claims, err := s.getUserFromRequest(r)
	if err != nil {
		logger.Error(
			"Failed to get user from request",
			slog.Any("error", err),
		)
		endpoint.Unauthorized(w)
		return
	}

	req := user.AcceptMembershipRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil && err != io.EOF {
		logger.Error(
			"Failed to decode request",
			slog.Any("error", err),
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	membership, token, err := s.acceptMembership(
		claims.AppId,
		usr.Id,
		organizationId,
		claims.Device,
		req.Activate,
	)
	if err != nil {
		if errors.Is(err, ErrPendingMembershipNotFound) {
			endpoint.Unauthorized(w)
			return
		}

		logger.Error(
			"Failed to accept membership",
			slog.Any("error", err),
			slog.String("organizationId", organizationId),
			slog.String("userId", usr.Id),
		)
		endpoint.InternalServerError(w)
		return
	}

	endpoint.WriteJSON(w, http.StatusOK, user.AcceptMembershipResponse{
		Membership: *membership,
		Token:      token,
	})
}

func (s *UserService) acceptMembership(
	appId string,
	userId string,
	organizationId string,
	device string,
	activate bool,
) (*user.Membership, *user.Token, error) {
	membership, found, err := s.findMembership(appId, userId, organizationId)
	if err != nil {
		return nil, nil, err
	}
	if !found || membership.Status != user.MembershipStatusPending {
		return nil, nil, ErrPendingMembershipNotFound
	}

	acceptedAt := time.Now()
	membership, _, err = s.updateMembership(
		membership,
		user.MembershipStatusAccepted,
		membership.Roles,
		membership.InvitedBy,
		&acceptedAt,
	)
	if err != nil {
		return nil, nil, err
	}

	if activate {
		if err := s.setActivation(appId, userId, device, organizationId); err != nil {
			return nil, nil, err
		}
	}

	if err := s.inactivateTokens(appId, userId); err != nil {
		return nil, nil, err
	}

	if !activate {
		return membership, nil, nil
	}

	userI, found, err := s.userStore.Query(datastore.Id(userId)).FindOne()
	if err != nil {
		return nil, nil, err
	}
	if !found {
		return nil, nil, errors.New("user not found")
	}

	token, err := s.issueToken(appId, userI.(*user.User), device)
	if err != nil {
		return nil, nil, err
	}

	return membership, token, nil
}
