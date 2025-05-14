package policyservice

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	policy "github.com/1backend/1backend/server/internal/services/policy/types"
	"golang.org/x/time/rate"
)

// @ID check
// @Summary Check
// @Description Check records a resource access and returns if the access is allowed.
// @Tags Policy Svc
// @Accept json
// @Produce json
// @Param body body policy.CheckRequest true "Check Request"
// @Success 200 {object} policy.CheckResponse "Checked successfully"
// @Failure 400 {object} policy.ErrorResponse "Invalid JSON"
// @Failure 401 {object} policy.ErrorResponse "Unauthorized"
// @Failure 500 {object} policy.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /policy-svc/check [post]
func (s *PolicyService) Check(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, statusCode, err := s.permissionChecker.HasPermission(
		r,
		policy.PermissionTemplateEdit,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.GetAuthorized() {
		endpoint.Unauthorized(w)
		return
	}

	req := policy.CheckRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error(
			"Error decoding request",
			slog.Any("error", err),
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	allowed, err := s.check(&req)
	if err != nil {
		logger.Error(
			"Error checking policy",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	bs, _ := json.Marshal(policy.CheckResponse{
		Allowed: allowed,
	})
	w.Write(bs)
}

func (s *PolicyService) check(request *policy.CheckRequest) (bool, error) {
	for _, instance := range s.instances {

		switch string(instance.TemplateId) {
		case policy.RateLimitPolicyTemplate.GetId():

			maxRequests := instance.Parameters.RateLimit.MaxRequests
			timeWindow, err := time.ParseDuration(
				instance.Parameters.RateLimit.TimeWindow,
			)
			if err != nil {
				return false, err
			}

			var limiterKey string
			switch instance.Parameters.RateLimit.Entity {
			case policy.EntityUserID:
				limiterKey = request.UserId
			case policy.EntityIP:
				limiterKey = request.Ip
			default:
				return false, fmt.Errorf("unknown entity type")
			}

			if instance.Parameters.RateLimit.Scope == policy.ScopeEndpoint {
				limiterKey += ":" + request.Endpoint
			}

			s.mutex.Lock()
			limiter, exists := s.rateLimiters.Load(limiterKey)
			if !exists {
				limiter = rate.NewLimiter(rate.Every(timeWindow), maxRequests)
				s.rateLimiters.Store(limiterKey, limiter)
			}
			s.mutex.Unlock()

			if !limiter.(*rate.Limiter).Allow() {
				return false, nil
			}
		}
	}

	return true, nil
}
