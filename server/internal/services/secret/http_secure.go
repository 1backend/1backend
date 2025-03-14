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
	"net/http"

	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	secret "github.com/openorch/openorch/server/internal/services/secret/types"
)

// @Id isSecure
// @Summary Check Security Status
// @Description Returns true if the encryption key is sufficiently secure.
// @Tags Secret Svc
// @Accept json
// @Produce json
// @Success 200 {object} secret.IsSecureResponse "Encrypt Value Response"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /secret-svc/is-secure [get]
func (cs *SecretService) Secure(
	w http.ResponseWriter,
	r *http.Request,
) {
	isAuthRsp, _, err := cs.clientFactory.Client(sdk.WithTokenFromRequest(r)).
		UserSvcAPI.IsAuthorized(r.Context(), *secret.PermissionSecretSave.Id).
		Body(openapi.UserSvcIsAuthorizedRequest{
			GrantedSlugs: []string{"model-svc"},
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

	jsonData, _ := json.Marshal(secret.IsSecureResponse{
		IsSecure: cs.encryptionKey != "changeMeToSomethingSecureForReal",
	})
	w.Write(jsonData)
}
