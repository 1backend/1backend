/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package container_svc

type BuildImageRequest struct {
	// Name is the name of the image to build
	Name string `json:"name" example:"nginx:latest" binding:"required"`

	// ContextPath is the local path to the build context
	ContextPath string `json:"contextPath" example:"." binding:"required"`

	// DockerfilePath is the local path to the Dockerfile
	DockerfilePath string `json:"dockerfilePath" example:"Dockerfile"`
}

type BuildImageResponse struct {
}

type ImagePullableRequest struct {
	ImageTag string `json:"imageTag" example:"crufter/llama-cpp-python:cuda-12.8.0-latest" binding:"required"`
}

type ImagePullableResponse struct {
	Pullable bool `json:"pullable" binding:"required"`
}
