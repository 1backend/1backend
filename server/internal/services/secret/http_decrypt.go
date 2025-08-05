/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package secretservice

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	"github.com/1backend/1backend/sdk/go/secrets"
	secret "github.com/1backend/1backend/server/internal/services/secret/types"
)

// @Id decryptValue
// @Summary Decrypt a Value
// @Description Decrypt a value and return the encrypted result
// @Tags Secret Svc
// @Accept json
// @Produce json
// @Param body body secret.DecryptValueRequest true "Decrypt Value Request"
// @Success 200 {object} secret.DecryptValueResponse "Decrypt Value Response"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /secret-svc/decrypt [post]
func (cs *SecretService) Decrypt(
	w http.ResponseWriter,
	r *http.Request,
) {
	isAuthRsp, statusCode, err := cs.options.PermissionChecker.HasPermission(
		r,
		secret.PermissionSecretSave,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.Authorized {
		endpoint.Unauthorized(w)
		return
	}

	req := &secret.DecryptValueRequest{}
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

	decryptedValue, err := secrets.Decrypt(req.Value, cs.options.SecretEncryptionKey)
	if err != nil {
		logger.Error(
			"Failed to decrypt value",
			slog.String("error", err.Error()),
			slog.String("value", req.Value),
		)
		endpoint.InternalServerError(w)
		return
	}

	jsonData, _ := json.Marshal(secret.DecryptValueResponse{
		Value: decryptedValue,
	})
	_, err = w.Write([]byte(jsonData))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}
