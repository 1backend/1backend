package userservice

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
)

// @ID listMemberships
// @Summary List Memberships
// @Description Lists organization memberships and pending invites.
// @Tags User Svc
// @Accept json
// @Produce json
// @Param body body user.ListMembershipsRequest false "List Memberships Request"
// @Success 200 {object} user.ListMembershipsResponse
// @Failure 400 {object} user.ErrorResponse "Invalid JSON"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Failure 500 {object} user.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /user-svc/memberships [post]
func (s *UserService) ListMemberships(
	w http.ResponseWriter,
	r *http.Request,
) {
	usr, claims, err := s.getUserFromRequest(r)
	if err != nil {
		logger.Error(
			"Failed to get user from request",
			slog.Any("error", err),
		)
		endpoint.Unauthorized(w)
		return
	}

	req := user.ListMembershipsRequest{}
	if r.Body != nil {
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
	}

	rsp, err := s.listMemberships(claims.AppId, usr.Id, &req)
	if err != nil {
		if err.Error() == "unauthorized" {
			endpoint.Unauthorized(w)
			return
		}

		logger.Error(
			"Failed to list memberships",
			slog.Any("error", err),
			slog.String("userId", usr.Id),
		)
		endpoint.InternalServerError(w)
		return
	}

	endpoint.WriteJSON(w, http.StatusOK, rsp)
}

func (s *UserService) listMemberships(
	appId string,
	callerUserId string,
	request *user.ListMembershipsRequest,
) (*user.ListMembershipsResponse, error) {
	roleIds, err := s.getRolesByUserId(appId, callerUserId)
	if err != nil {
		return nil, err
	}

	isPlatformAdmin := contains(roleIds, user.RoleAdmin)
	canListOrganization := request.OrganizationId != "" && (isPlatformAdmin || hasOrganizationAdminAccess(roleIds, request.OrganizationId))

	effectiveUserId := request.UserId
	if effectiveUserId == "" {
		if canListOrganization {
			effectiveUserId = ""
		} else {
			effectiveUserId = callerUserId
		}
	}

	if effectiveUserId != "" && effectiveUserId != callerUserId && !isPlatformAdmin && !canListOrganization {
		return nil, fmt.Errorf("unauthorized")
	}

	filters := []datastore.Filter{
		datastore.Equals(datastore.Field("appId"), appId),
	}
	if request.OrganizationId != "" {
		filters = append(filters, datastore.Equals(datastore.Field("organizationId"), request.OrganizationId))
	}
	if effectiveUserId != "" {
		filters = append(filters, datastore.Equals(datastore.Field("userId"), effectiveUserId))
	}
	if request.Status != "" {
		filters = append(filters, datastore.Equals(datastore.Field("status"), request.Status))
	}

	q := s.membershipStore.Query(filters...).OrderBy(
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

	membershipIs, err := q.Find()
	if err != nil {
		return nil, err
	}
	if len(membershipIs) == 0 {
		return &user.ListMembershipsResponse{
			Memberships: []struct {
				Membership   user.Membership   `json:"membership" binding:"required"`
				Organization user.Organization `json:"organization" binding:"required"`
				User         user.User         `json:"user" binding:"required"`
			}{},
		}, nil
	}

	organizationIds := []any{}
	userIds := []any{}
	organizationMap := map[string]struct{}{}
	userMap := map[string]struct{}{}
	for _, membershipI := range membershipIs {
		membership := membershipI.(*user.Membership)
		if _, ok := organizationMap[membership.OrganizationId]; !ok {
			organizationIds = append(organizationIds, membership.OrganizationId)
			organizationMap[membership.OrganizationId] = struct{}{}
		}
		if _, ok := userMap[membership.UserId]; !ok {
			userIds = append(userIds, membership.UserId)
			userMap[membership.UserId] = struct{}{}
		}
	}

	orgIs, err := s.organizationStore.Query(
		datastore.Equals(datastore.Field("appId"), appId),
		datastore.IsInList(datastore.Field("id"), organizationIds...),
	).Find()
	if err != nil {
		return nil, err
	}

	userIs, err := s.userStore.Query(
		datastore.IsInList(datastore.Field("id"), userIds...),
	).Find()
	if err != nil {
		return nil, err
	}

	orgsById := map[string]*user.Organization{}
	for _, orgI := range orgIs {
		org := orgI.(*user.Organization)
		orgsById[org.Id] = org
	}

	usersById := map[string]*user.User{}
	for _, userI := range userIs {
		usr := userI.(*user.User)
		usersById[usr.Id] = usr
	}

	records := make([]struct {
		Membership   user.Membership   `json:"membership" binding:"required"`
		Organization user.Organization `json:"organization" binding:"required"`
		User         user.User         `json:"user" binding:"required"`
	}, 0, len(membershipIs))
	for _, membershipI := range membershipIs {
		membership := membershipI.(*user.Membership)
		org := orgsById[membership.OrganizationId]
		usr := usersById[membership.UserId]
		if org == nil || usr == nil {
			return nil, fmt.Errorf("membership references missing user or organization")
		}
		records = append(records, struct {
			Membership   user.Membership   `json:"membership" binding:"required"`
			Organization user.Organization `json:"organization" binding:"required"`
			User         user.User         `json:"user" binding:"required"`
		}{
			Membership:   *membership,
			Organization: *org,
			User:         *usr,
		})
	}

	return &user.ListMembershipsResponse{
		Memberships: records,
	}, nil
}
