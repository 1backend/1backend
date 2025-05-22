package containerservice

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
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

	isAuthRsp, statusCode, err := dm.options.PermissionChecker.HasPermission(
		r,
		container.PermissionImageBuild,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.GetAuthorized() {
		endpoint.Unauthorized(w)
		return
	}

	req := &container.BuildImageRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	_, err = dm.backend.BuildImage(*req)
	if err != nil {
		logger.Error("Failed to build image",
			slog.String("dockerfilePath", req.DockerfilePath),
			slog.String("contextPath", req.ContextPath),
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	json.NewEncoder(w).Encode(&container.BuildImageResponse{})
}
