/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package containerservice

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	container "github.com/1backend/1backend/server/internal/services/container/types"
	"github.com/gorilla/mux"
)

// @ID imagePullable
// @Summary      Check if Container Image is Pullable
// @Description  Check if an image exists on in the container registry and is pullable.
// @Tags         Container Svc
// @Accept       json
// @Produce      json
// @Param imageName path string true "Image name"
// @Success      200   {object}  container.ImagePullableResponse
// @Failure      400   {object}  container.ErrorResponse  "model ID in path is not URL encoded"
// @Failure      401   {object}  container.ErrorResponse  "unauthorized"
// @Failure      500   {object}  container.ErrorResponse  "internal server error"
// @Security BearerAuth
// @Router       /container-svc/image/{imageName}/pullable [get]
func (dm *ContainerService) ImagePullable(
	w http.ResponseWriter,
	req *http.Request,
) {

	isAuthRsp, statusCode, err := dm.options.PermissionChecker.HasPermission(
		req,
		container.PermissionContainerView,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.GetAuthorized() {
		endpoint.Unauthorized(w)
		return
	}

	vars := mux.Vars(req)
	imageName, err := url.PathUnescape(vars["imageName"])
	if err != nil {
		logger.Error(
			"Failed to unescape image name",
			"error", err,
			"imageName", vars["imageName"],
		)
		endpoint.WriteString(w, http.StatusBadRequest, "model ID in path is not URL encoded")
		return
	}

	pullable, err := dm.imagePullable(imageName)
	if err != nil {
		logger.Error(
			"Failed to check if image is pullable",
			"error", err,
			"imageName", imageName,
		)
		endpoint.InternalServerError(w)
		return
	}

	jsonData, _ := json.Marshal(container.ImagePullableResponse{
		Pullable: pullable,
	})
	_, err = w.Write([]byte(jsonData))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

type DockerHubResponse struct {
	Name      string `json:"name"`
	TagStatus string `json:"tag_status"`
}

func (dm *ContainerService) imagePullable(imageName string) (bool, error) {
	parts := strings.Split(imageName, ":")
	image := parts[0]
	tag := "latest"
	if len(parts) > 1 {
		tag = parts[1]
	}

	return imageExistsInRegistry(image, tag)
}

func imageExistsInRegistry(image, tag string) (bool, error) {
	if !strings.Contains(image, "/") {
		image = "library/" + image
	}

	url := fmt.Sprintf(
		"https://hub.docker.com/v2/repositories/%s/tags/%s",
		image,
		url.PathEscape(tag),
	)

	resp, err := http.Get(url)
	if err != nil {
		return false, fmt.Errorf("error making request to Docker Hub: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return false, nil
	}

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("unexpected response from Docker Hub: %v", resp.Status)
	}

	return true, nil

	// @crufter images that exist but haven't been pulled for a while seem to be marked as inactive
	//
	// var response DockerHubResponse
	// err = json.NewDecoder(resp.Body).Decode(&response)
	// if err != nil {
	// 	return false, fmt.Errorf("error decoding response: %v", err)
	// }
	//
	// return response.TagStatus == "active", nil
}
