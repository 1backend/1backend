package userservice

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

// @ID declineMembership
// @Summary Decline Membership
// @Description Declines the caller user's pending invite for an organization.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param organizationId path string true "Organization ID"
// @Param body body user.DeclineMembershipRequest false "Decline Membership Request"
// @Success 200 {object} user.DeclineMembershipResponse
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/organization/{organizationId}/membership/decline [post]
func (s *UserService) DeclineMembership(
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

	req := user.DeclineMembershipRequest{}
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

	membership, err := s.declineMembership(
		claims.AppId,
		usr.Id,
		organizationId,
	)
	if err != nil {
		if errors.Is(err, ErrPendingMembershipNotFound) {
			endpoint.Unauthorized(w)
			return
		}

		logger.Error(
			"Failed to decline membership",
			slog.Any("error", err),
			slog.String("organizationId", organizationId),
			slog.String("userId", usr.Id),
		)
		endpoint.InternalServerError(w)
		return
	}

	endpoint.WriteJSON(w, http.StatusOK, user.DeclineMembershipResponse{
		Membership: *membership,
	})
}

func (s *UserService) declineMembership(
	appId string,
	userId string,
	organizationId string,
) (*user.Membership, error) {
	membership, found, err := s.findMembership(appId, userId, organizationId)
	if err != nil {
		return nil, err
	}
	if !found || membership.Status != user.MembershipStatusPending {
		return nil, ErrPendingMembershipNotFound
	}

	membership, _, err = s.updateMembership(
		membership,
		user.MembershipStatusDeclined,
		membership.Roles,
		membership.InvitedBy,
		nil,
	)
	if err != nil {
		return nil, err
	}

	if err := s.deleteActivations(appId, userId, organizationId); err != nil {
		return nil, err
	}

	if err := s.inactivateTokens(appId, userId); err != nil {
		return nil, err
	}

	return membership, nil
}
