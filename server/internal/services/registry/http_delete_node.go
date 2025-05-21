package registryservice

import (
	"log/slog"
	"net/http"
	"net/url"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	registry "github.com/1backend/1backend/server/internal/services/registry/types"
	"github.com/gorilla/mux"
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

	isAuthRsp, statusCode, err := rs.options.PermissionChecker.HasPermission(
		r,
		registry.PermissionNodeDelete,
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
	nodeURL, err := url.PathUnescape(vars["url"])
	if err != nil || nodeURL == "" {
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid URL")
		return
	}

	err = rs.deleteNodeByURL(nodeURL)
	if err != nil {
		if err == registry.ErrNotFound {
			endpoint.WriteString(w, http.StatusNotFound, "Not Found")
		} else {
			logger.Error("Error deleting node",
				slog.String("url", nodeURL),
				slog.Any("error", err),
			)
			endpoint.InternalServerError(w)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (rs *RegistryService) deleteNodeByURL(ur string) error {
	return rs.nodeStore.Query(datastore.Equals([]string{"url"}, ur)).Delete()
}
