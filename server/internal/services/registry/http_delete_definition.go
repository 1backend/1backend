package registryservice

import (
	"net/http"

	"github.com/gorilla/mux"
	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/datastore"
	registry "github.com/openorch/openorch/server/internal/services/registry/types"
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
// @Failure 404 {object} registry.ErrorResponse "Service not found"
// @Failure 500 {object} registry.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /registry-svc/definition/{id} [delete]
func (rs *RegistryService) DeleteDefinition(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, _, err := rs.clientFactory.Client(sdk.WithTokenFromRequest(r)).
		UserSvcAPI.IsAuthorized(r.Context(), *registry.PermissionDefinitionDelete.Id).
		Body(openapi.UserSvcIsAuthorizedRequest{
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

	vars := mux.Vars(r)
	serviceID := vars["id"]
	if serviceID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid definition ID`))
		return
	}

	err = rs.deleteDefinitionByID(serviceID)
	if err != nil {
		if err == registry.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`Definition not found`))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (rs *RegistryService) deleteDefinitionByID(id string) error {
	return rs.definitionStore.Query(datastore.Id(id)).Delete()
}
