package registryservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	openapi "github.com/1backend/1backend/clients/go"
	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/datastore"
	registry "github.com/1backend/1backend/server/internal/services/registry/types"
	"github.com/pkg/errors"
)

// @ID registerInstance
// @Summary Register Instance
// @Description Registers an instance. Idempotent.
// @Tags Registry Svc
// @Accept json
// @Produce json
// @Param body body registry.RegisterInstanceRequest true "Register Instance Request"
// @Success 201 {object} registry.RegisterInstanceResponse
// @Failure 400 {object} registry.ErrorResponse "Invalid JSON"
// @Failure 401 {object} registry.ErrorResponse "Unauthorized"
// @Failure 500 {object} registry.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /registry-svc/instance [put]
func (rs *RegistryService) RegisterInstance(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, _, err := rs.clientFactory.Client(client.WithTokenFromRequest(r)).
		UserSvcAPI.HasPermission(r.Context(), registry.PermissionInstanceEdit).
		Body(openapi.UserSvcHasPermissionRequest{
			GrantedSlugs: []string{"deploy-svc"},
		}).
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

	req := &registry.RegisterInstanceRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	err = rs.registerInstance(isAuthRsp.User.Slug, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	go func() {
		rs.triggerChan <- struct{}{}
		// @todo remove this and the tests fail ???
		rs.triggerChan <- struct{}{}
	}()

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{}`))
}

func (rs *RegistryService) registerInstance(
	callerSlug string,
	req *registry.RegisterInstanceRequest,
) error {
	var instance registry.Instance

	instances, err := rs.instanceStore.Query(datastore.Equals([]string{"url"}, req.URL)).
		Find()
	if err != nil {
		return err
	}

	if len(instances) > 0 {
		instance = *instances[0].(*registry.Instance)
	} else {
		instance.Id = sdk.Id("inst")
	}

	nu := len(instances) == 0
	if nu {
		instance.Slug = callerSlug
	}

	if callerSlug != "deploy-svc" && callerSlug != instance.Slug {
		return fmt.Errorf(
			"caller slug '%s' does not match instance slug '%s'. perhaps a misconfigured service already took this URL. if you are an admin you can fix that by calling 'oo instance remove %v'",
			callerSlug,
			instance.Slug,
			instance.Id,
		)
	}

	if req.Id != "" {
		instance.Id = req.Id
	}
	if req.URL != "" {
		instance.URL = req.URL
	}
	if req.Scheme != "" {
		instance.Scheme = req.Scheme
	}
	if req.Host != "" {
		instance.Host = req.Host
	}
	if req.IP != "" {
		instance.IP = req.IP
	}
	if req.Path != "" {
		instance.Path = req.Path
	}
	if req.DeploymentId != "" {
		instance.DeploymentId = req.DeploymentId
	}

	if instance.Status == "" {
		instance.Status = registry.InstanceStatusUnknown
	}

	if instance.URL == "" {
		return errors.New("url is missing")
	}

	return rs.instanceStore.Upsert(&instance)
}
