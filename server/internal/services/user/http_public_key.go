/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/1backend/1backend/sdk/go/logger"
	user "github.com/1backend/1backend/server/internal/services/user/types"
)

// @ID getPublicKey
// @Summary Get Public Key
// @Description Get the public key to verify the JWT signature.
// @Tags User Svc
// @Accept json
// @Produce json
// @Success 200 {object} user.GetPublicKeyResponse
// @Failure 400 {object} user.ErrorResponse "Invalid JSON or missing permission id"
// @Failure 401 {object} user.ErrorResponse "Unauthorized"
// @Router /user-svc/public-key [get]
func (s *UserService) GetPublicKey(
	w http.ResponseWriter,
	r *http.Request) {

	bs, _ := json.Marshal(user.GetPublicKeyResponse{
		PublicKey: s.publicKeyPem,
	})
	_, err := w.Write(bs)
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}
