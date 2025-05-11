package registryservice

import (
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	registry "github.com/1backend/1backend/server/internal/services/registry/types"
	"github.com/gorilla/mux"
)

// @ID deleteDefinition
// @Summary Delete Definition
// @Description Deletes a registered definition by ID.
// @Tags Registry Svc
// @Accept json
// @Produce json
// @Param id path string true "Definition ID"
// @Success 204 "No Content"
// @Failure 400 {object} registry.ErrorResponse "Invalid ID"
// @Failure 401 {object} registry.ErrorResponse "Unauthorized"
// @Failure 404 {object} registry.ErrorResponse "Not Found"
// @Failure 500 {object} registry.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /registry-svc/definition/{id} [delete]
func (rs *RegistryService) DeleteDefinition(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, statusCode, err := rs.permissionChecker.HasPermission(
		r,
		registry.PermissionDefinitionDelete,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.GetAuthorized() {
		endpoint.Unauthorized(w)
		return
	}

	vars := mux.Vars(r)
	serviceID := vars["id"]
	if serviceID == "" {
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	err = rs.deleteDefinitionByID(serviceID)
	if err != nil {
		if err == registry.ErrNotFound {
			endpoint.WriteString(w, http.StatusNotFound, "Not Found")
		} else {
			logger.Error("Error deleting definition",
				slog.String("id", serviceID),
				slog.Any("error", err),
			)
			endpoint.InternalServerError(w)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (rs *RegistryService) deleteDefinitionByID(id string) error {
	return rs.definitionStore.Query(datastore.Id(id)).Delete()
}
