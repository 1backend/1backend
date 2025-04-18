package containerservice

import (
	"encoding/json"
	"net/http"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	container "github.com/1backend/1backend/server/internal/services/container/types"
)

// @ID buildImage
// @Summary Build an Image
// @Description Builds a Docker image with the specified parameters.
// @Description
// @Description Requires the `container-svc:image:build` permission.
// @Tags Container Svc
// @Accept json
// @Produce json
// @Param body body container.BuildImageRequest true "Build Image Request"
// @Success 200 {object} container.BuildImageResponse
// @Failure 400 {object} container.ErrorResponse "Invalid JSON"
// @Failure 401 {object} container.ErrorResponse "Unauthorized"
// @Failure 500 {object} container.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /container-svc/image [put]
func (dm *ContainerService) BuildImage(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, _, err := dm.clientFactory.Client(client.WithTokenFromRequest(r)).
		UserSvcAPI.HasPermission(r.Context(), container.PermissionImageBuild).
		Body(openapi.UserSvcHasPermissionRequest{
			PermittedSlugs: []string{"deploy-svc"},
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

	req := &container.BuildImageRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(container.ErrorResponse{Error: "Invalid JSON"})
		return
	}
	defer r.Body.Close()

	_, err = dm.backend.BuildImage(*req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&container.BuildImageResponse{})
}
