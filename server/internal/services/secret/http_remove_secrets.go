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
	"context"
	"encoding/json"
	"net/http"

	openapi "github.com/openorch/openorch/clients/go"
	sdk "github.com/openorch/openorch/sdk/go"
	"github.com/openorch/openorch/sdk/go/datastore"
	secret "github.com/openorch/openorch/server/internal/services/secret/types"
	"github.com/pkg/errors"
)

// @Id removeSecrets
// @Summary Remove Secrets
// @Description Remove secrets if authorized to do so
// @Tags Secret Svc
// @Accept json
// @Produce json
// @Param body body secret.RemoveSecretsRequest true "Remove Secret Request"
// @Success 200 {object} secret.RemoveSecretsResponse "Remove Secret Response"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /secret-svc/secrets [delete]
func (cs *SecretService) RemoveSecrets(
	w http.ResponseWriter,
	r *http.Request,
) {
	isAuthRsp, _, err := cs.clientFactory.Client(sdk.WithTokenFromRequest(r)).
		UserSvcAPI.IsAuthorized(r.Context(), *secret.PermissionSecretRemove.Id).
		Body(openapi.UserSvcIsAuthorizedRequest{}).
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

	req := &secret.RemoveSecretsRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
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

	err = cs.removeSecrets(
		r.Context(),
		req,
		isAdmin,
		*isAuthRsp.User.Slug,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(secret.RemoveSecretsResponse{})
	w.Write(jsonData)
}

func (cs *SecretService) removeSecrets(
	ctx context.Context,
	req *secret.RemoveSecretsRequest,
	isAdmin bool,
	userSlug string,
) error {
	cs.lock.Acquire(ctx, "secret-svc-save")
	defer cs.lock.Release(ctx, "secret-svc-save")

	keys := []any{}
	for _, key := range req.Keys {
		keys = append(keys, key)
	}
	if req.Key != "" {
		keys = append(keys, req.Key)
	}

	ids := []any{}
	for _, id := range req.Ids {
		ids = append(ids, id)
	}

	if req.Id != "" {
		ids = append(ids, req.Id)
	}

	secrets := []*secret.Secret{}

	secretIs, err := cs.secretStore.Query(datastore.IsInList([]string{"key"}, keys...)).Find()
	if err != nil {
		return errors.Wrap(err, "failed to query secrets")
	}
	for _, secretI := range secretIs {
		s := secretI.(*secret.Secret)
		secrets = append(secrets, s)
	}

	secretIs, err = cs.secretStore.Query(datastore.IsInList([]string{"id"}, ids...)).Find()
	if err != nil {
		return errors.Wrap(err, "failed to query secrets")
	}
	for _, secretI := range secretIs {
		s := secretI.(*secret.Secret)
		secrets = append(secrets, s)
	}

	for _, s := range secrets {
		canRemove := isAdmin
		if !canRemove {
			for _, deleter := range s.Deleters {
				if deleter == userSlug {
					canRemove = true
					break
				}
			}
		}
		if !canRemove {
			continue
		}

		err = cs.secretStore.Query(datastore.Equals([]string{"id"}, s.Id)).Delete()
		if err != nil {
			return errors.Wrap(err, "failed to delete secret")
		}
	}

	return nil
}
