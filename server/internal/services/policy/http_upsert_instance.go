package policyservice

import (
	"encoding/json"
	"log/slog"
	"net/http"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	policy "github.com/1backend/1backend/server/internal/services/policy/types"
	"github.com/gorilla/mux"
)

// @ID upsertInstance
// @Summary Upsert an Instance
// @Description Allows user to upsert a new policy instance based on a template.
// @Tags Policy Svc
// @Accept json
// @Produce json
// @Param instanceId path string true "Instance ID"
// @Param body body policy.UpsertInstanceRequest true "Upsert Instance Request"
// @Success 200 {object} policy.UpsertInstanceResponse "Instance upserted successfully"
// @Failure 400 {object} policy.ErrorResponse "Invalid JSON"
// @Failure 401 {object} policy.ErrorResponse "Unauthorized"
// @Failure 500 {object} policy.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /policy-svc/instance/{instanceId} [put]
func (s *PolicyService) UpsertInstance(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, statusCode, err := s.options.PermissionChecker.HasPermission(
		r,
		policy.PermissionInstanceEdit,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.Authorized {
		endpoint.Unauthorized(w)
		return
	}

	req := policy.UpsertInstanceRequest{}
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

	req.Instance.Id = mux.Vars(r)["instanceId"]

	err = s.upsertInstance(req.Instance)
	if err != nil {
		logger.Error("Error upserting instance", slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	bs, _ := json.Marshal(policy.UpsertInstanceResponse{})
	_, err = w.Write(bs)
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func (s *PolicyService) upsertInstance(instance *policy.Instance) error {
	if instance.Id == "" {
		instance.Id = sdk.Id("insta")
	}

	exists := false
	for _, i := range s.instances {
		if i.Id == instance.Id {
			exists = true
		}
	}
	if exists {
		return nil
	}
	s.instances = append(s.instances, instance)
	return s.instancesStore.Upsert(instance)
}
