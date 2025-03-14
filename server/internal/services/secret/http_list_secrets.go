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

	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/datastore"
	secret "github.com/openorch/openorch/server/internal/services/secret/types"
	"github.com/pkg/errors"
)

// @Id listSecrets
// @Summary List Secrets
// @Description List secrets by key(s) if authorized.
// @Tags Secret Svc
// @Accept json
// @Produce json
// @Param body body secret.ListSecretsRequest false "List Secret Request"
// @Success 200 {object} secret.ListSecretsResponse "List Secret Response"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /secret-svc/secrets [post]
func (cs *SecretService) ListSecrets(
	w http.ResponseWriter,
	r *http.Request,
) {
	isAuthRsp, _, err := cs.clientFactory.Client(sdk.WithTokenFromRequest(r)).
		UserSvcAPI.IsAuthorized(r.Context(), *secret.PermissionSecretList.Id).
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

	req := secret.ListSecretsRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid JSON`))
		return
	}
	defer r.Body.Close()

	isAdmin, err := cs.authorizer.IsAdminFromRequest(cs.publicKey, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	ss, err := cs.getSecrets(req, isAdmin, *isAuthRsp.User.Slug)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(secret.ListSecretsResponse{
		Secrets: ss,
	})
	w.Write(jsonData)
}

func (cs *SecretService) getSecrets(
	req secret.ListSecretsRequest, isAdmin bool, userSlug string,
) ([]*secret.Secret, error) {
	filters := []datastore.Filter{}
	if req.Namespace == "" {
		req.Namespace = "default"
	}
	filters = append(filters, datastore.Equals([]string{"namespace"}, req.Namespace))

	if req.Key != "" {
		filters = append(filters, datastore.Equals([]string{"key"}, req.Key))
	}
	if req.Keys != nil {
		keys := []any{}
		for _, v := range req.Keys {
			keys = append(keys, v)
		}
		filters = append(filters, datastore.IsInList([]string{"key"}, keys...))
	}

	secretIs, err := cs.secretStore.Query(filters...).Find()
	if err != nil {
		return nil, err
	}

	secrets := []*secret.Secret{}
	for _, secretI := range secretIs {
		s := secretI.(*secret.Secret)
		canList := isAdmin

		if !canList {
			for _, slug := range s.Readers {
				if slug == userSlug {
					canList = true
					break
				}
			}
		}

		if !canList {
			continue
		}

		s.Value, err = decrypt(s.Value, cs.encryptionKey)
		if err != nil {
			return nil, errors.Wrap(err, "failed to decrypt secret")
		}

		secrets = append(secrets, s)
	}

	return secrets, nil
}
