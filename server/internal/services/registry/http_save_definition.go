package registryservice

import (
	"encoding/json"
	"net/http"

	sdk "github.com/openorch/openorch/sdk/go"
	registry "github.com/openorch/openorch/server/internal/services/registry/types"
)

// Register a new definition
// @ID saveDefinition
// @Summary Register a Definition
// @Description Registers a new definition, associating an definition address with a slug acquired from the bearer token.
// @Tags Registry Svc
// @Accept json
// @Produce json
// @Param body body registry.SaveDefinitionRequest true "Register Service Definition Request"
// @Success 201 {object} registry.SaveDefinitionResponse
// @Failure 400 {object} registry.ErrorResponse "Invalid JSON"
// @Failure 401 {object} registry.ErrorResponse "Unauthorized"
// @Failure 500 {object} registry.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /registry-svc/definition [put]
func (rs *RegistryService) SaveDefinition(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, _, err := rs.clientFactory.Client(sdk.WithTokenFromRequest(r)).
		UserSvcAPI.IsAuthorized(r.Context(), *registry.PermissionNodeView.Id).
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

	req := &registry.SaveDefinitionRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	err = rs.saveServiceDefinition(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{}`))
}

func (rs *RegistryService) saveServiceDefinition(
	req *registry.SaveDefinitionRequest,
) error {
	return rs.definitionStore.Upsert(req.Definition)
}
