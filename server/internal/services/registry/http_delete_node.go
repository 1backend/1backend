package registryservice

import (
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/datastore"
	registry "github.com/openorch/openorch/server/internal/services/registry/types"
)

// @ID deleteNode
// @Summary Delete Node
// @Description Deletes a registered node by node URL. This endpoint is useful when a node is no longer available but it's still present in the database.
// @Tags Registry Svc
// @Accept json
// @Produce json
// @Param url path string true "Node URL"
// @Success 204 "No Content"
// @Failure 400 {object} registry.ErrorResponse "Invalid ID"
// @Failure 401 {object} registry.ErrorResponse "Unauthorized"
// @Failure 404 {object} registry.ErrorResponse "Service not found"
// @Failure 500 {object} registry.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /registry-svc/node/{url} [delete]
func (rs *RegistryService) DeleteNode(
	w http.ResponseWriter,
	r *http.Request,
) {

	isAuthRsp, _, err := rs.clientFactory.Client(sdk.WithTokenFromRequest(r)).
		UserSvcAPI.IsAuthorized(r.Context(), *registry.PermissionNodeDelete.Id).
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
	nodeURL, err := url.PathUnescape(vars["url"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if nodeURL == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid node ID`))
		return
	}

	err = rs.deleteNodeByURL(nodeURL)
	if err != nil {
		if err == registry.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`Service not found`))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (rs *RegistryService) deleteNodeByURL(ur string) error {
	return rs.nodeStore.Query(datastore.Equals([]string{"url"}, ur)).Delete()
}
