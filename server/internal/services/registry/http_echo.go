/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package registryservice

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
)

// @ID echoPut
// @Summary Echo the request body in the response body.
// @Description This endpoint is used to test the server's response to a request.
// @Description It simply echoes back the request body as a JSON response.
// @Tags Registry Svc
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Echoed response"
// @Failure 400 {object} map[string]string "Invalid JSON"
// @Security BearerAuth
// @Router /registry-svc/echo [put]
func (ns *RegistryService) EchoPut(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close()

	var body interface{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		logger.Error(
			"Error decoding request body",
			slog.Any("error", err),
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(body); err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		endpoint.InternalServerError(w)
	}
}

// @ID echoPost
// @Summary Echo the request body in the response body.
// @Description This endpoint is used to test the server's response to a request.
// @Description It simply echoes back the request body as a JSON response.
// @Tags Registry Svc
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Echoed response"
// @Failure 400 {object} map[string]string "Invalid JSON"
// @Security BearerAuth
// @Router /registry-svc/echo [post]
func (ns *RegistryService) EchoPost(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close()

	var body interface{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		logger.Error(
			"Error decoding request body",
			slog.Any("error", err),
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(body); err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		endpoint.InternalServerError(w)
	}
}

// @ID echoGet
// @Summary Echo the query parameters in the response body.
// @Description This endpoint is used to test the server's response to a GET request.
// @Description It echoes back the query parameters as a JSON object.
// @Tags Registry Svc
// @Produce json
// @Success 200 {object} map[string]interface{} "Echoed query parameters"
// @Security BearerAuth
// @Router /registry-svc/echo [get]
func (ns *RegistryService) EchoGet(
	w http.ResponseWriter,
	r *http.Request,
) {
	query := r.URL.Query()
	response := make(map[string]interface{})

	for key, values := range query {
		if len(values) == 1 {
			response[key] = values[0]
		} else {
			response[key] = values
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		endpoint.InternalServerError(w)
	}
}
