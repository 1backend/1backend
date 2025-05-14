package registryservice

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/auth"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	registry "github.com/1backend/1backend/server/internal/services/registry/types"
	"github.com/samber/lo"
)

var permittedSlugs = []string{
	"deploy-svc",
	"proxy-svc",
}

// @ID listInstances
// @Summary List Service Instances
// @Description Retrieves a list of all instances or filters them by specific criteria (e.g., host, IP).
// @Tags Registry Svc
// @Accept json
// @Produce json
// @Param scheme query string false "Scheme to filter by"
// @Param ip query string false "IP to filter by"
// @Param deploymentId query string false "Deployment ID to filter by"
// @Param host query string false "Host to filter by"
// @Param ip query string false "IP to filter by"
// @Param id query string false "Id to filter by"
// @Param slug query string false "Slug to filter by"
// @Success 200 {object} registry.ListInstancesResponse
// @Failure 400 {object} registry.ErrorResponse "Invalid filters"
// @Failure 500 {object} registry.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /registry-svc/instances [get]
func (rs *RegistryService) ListInstances(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, statusCode, err := rs.permissionChecker.HasPermission(
		r,
		registry.PermissionInstanceView,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.GetAuthorized() {
		endpoint.Unauthorized(w)
		return
	}

	q := r.URL.Query()
	host := q.Get("host")
	ip := q.Get("ip")
	deploymentId := q.Get("deploymentId")
	id := q.Get("ip")
	path := q.Get("path")
	slug := q.Get("slug")

	publicKey, err := rs.getPublicKey()
	if err != nil {
		endpoint.WriteErr(w, http.StatusInternalServerError, err)
		return
	}

	isAdmin, err := auth.AuthorizerImpl{}.IsAdminFromRequest(publicKey, r)
	if err != nil {
		logger.Error(
			"Error checking if user is admin",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	instances, err := rs.getInstances(isAuthRsp.User.Slug, isAdmin, List{
		Id:           id,
		Host:         host,
		DeploymentId: deploymentId,
		IP:           ip,
		Path:         path,
		Slug:         slug,
	})
	if err != nil {
		logger.Error(
			"Error listing instances",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	bs, _ := json.Marshal(registry.ListInstancesResponse{
		Instances: instances,
	})
	_, err = w.Write(bs)
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

type List struct {
	Id           string
	DeploymentId string
	ServiceSlug  string
	Host         string
	IP           string
	Scheme       string
	Path         string
	Slug         string
}

func (rs *RegistryService) getInstances(
	callerSlug string,
	isAdmin bool,
	query List,
) ([]*registry.Instance, error) {
	instanceIs, err := rs.instanceStore.Query().Find()
	if err != nil {
		return nil, err
	}

	instances := []*registry.Instance{}
	for _, instanceI := range instanceIs {
		instances = append(instances, instanceI.(*registry.Instance))
	}

	filtered := []*registry.Instance{}
	for _, v := range instances {
		if !isAdmin &&
			v.Slug != callerSlug &&
			!lo.Contains(permittedSlugs, callerSlug) {
			continue
		}

		match := true

		if query.IP != "" && v.IP != query.IP {
			match = false
		}

		if query.Scheme != "" && v.Scheme != query.Scheme {
			match = false
		}

		if query.Host != "" && v.Host != query.Host {
			match = false
		}

		if query.Path != "" && v.Path != query.Path {
			match = false
		}
		if query.DeploymentId != "" && v.DeploymentId != query.DeploymentId {
			match = false
		}

		if query.Slug != "" && v.Slug != query.Slug {
			match = false
		}

		if match {
			filtered = append(filtered, v)
		}
	}

	return filtered, nil
}
