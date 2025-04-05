package policyservice

import (
	"encoding/json"
	"net/http"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/client"
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

	isAuthRsp, _, err := s.clientFactory.Client(client.WithTokenFromRequest(r)).
		UserSvcAPI.HasPermission(r.Context(), policy.PermissionInstanceEdit).
		Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if !isAuthRsp.GetAuthorized() {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`Unauthorized`))
		return
	}

	req := policy.UpsertInstanceRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	req.Instance.Id = mux.Vars(r)["instanceId"]

	err = s.upsertInstance(req.Instance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bs, _ := json.Marshal(policy.UpsertInstanceResponse{})
	w.Write(bs)
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
