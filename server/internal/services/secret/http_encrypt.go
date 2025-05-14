/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package secretservice

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	secret "github.com/1backend/1backend/server/internal/services/secret/types"
)

// @Id encryptValue
// @Summary Encrypt a Value
// @Description Encrypt a value and return the encrypted result
// @Tags Secret Svc
// @Accept json
// @Produce json
// @Param body body secret.EncryptValueRequest true "Encrypt Value Request"
// @Success 200 {object} secret.EncryptValueResponse "Encrypt Value Response"
// @Failure 400 {string} string "Invalid JSON"
// @Failure 400 {string} string "Missing Data"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /secret-svc/encrypt [post]
func (cs *SecretService) Encrypt(
	w http.ResponseWriter,
	r *http.Request,
) {
	isAuthRsp, statusCode, err := cs.permissionChecker.HasPermission(
		r,
		secret.PermissionSecretSave,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.GetAuthorized() {
		endpoint.Unauthorized(w)
		return
	}

	req := &secret.EncryptValueRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		logger.Error(
			"Failed to decode request body",
			slog.String("error", err.Error()),
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	if req.Value == "" && req.Values == nil {
		endpoint.WriteString(w, http.StatusBadRequest, "Missing Data")
		return
	}

	encryptedValue, err := encrypt(req.Value, cs.encryptionKey)
	if err != nil {
		logger.Error(
			"Failed to encrypt value",
			slog.String("error", err.Error()),
		)
		endpoint.InternalServerError(w)
		return
	}

	jsonData, _ := json.Marshal(secret.EncryptValueResponse{
		Value: encryptedValue,
	})
	_, err = w.Write([]byte(jsonData))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}
