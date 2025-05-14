package basicservice

import (
	"net/http"

	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"

	basic "github.com/1backend/1backend/examples/go/services/basic/internal/types"
)

// @ID error
// @Summary Error
// @Description This endpoint simply errors. Useful for testing the proxy.
// @Tags Basic Svc
// @Accept json
// @Produce json
// @Failure 500 {object} basic.ErrorResponse "Error"
// @Router /basic-svc/error [post]
func (s *BasicService) Error(w http.ResponseWriter, r *http.Request) {
	// We just use this value here so gofmt doesn't remove the unused basic package.
	logger.Info("Error", basic.ErrorResponse{})

	endpoint.InternalServerError(w)
}
