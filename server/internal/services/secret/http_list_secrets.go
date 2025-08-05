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

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	secretssdk "github.com/1backend/1backend/sdk/go/secrets"
	secret "github.com/1backend/1backend/server/internal/services/secret/types"
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

	isAuthRsp, statusCode, err := cs.options.PermissionChecker.HasPermission(
		r,
		secret.PermissionSecretList,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.Authorized {
		endpoint.Unauthorized(w)
		return
	}

	req := secret.ListSecretsRequest{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Error(
			"Failed to decode request",
			slog.Any("error", err),
		)
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()

	isAdmin, err := cs.options.Authorizer.IsAdminFromRequest(cs.publicKey, r)
	if err != nil {
		logger.Error(
			"Failed to check if user is admin",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	ss, err := cs.getSecrets(*isAuthRsp.App, req, isAdmin, isAuthRsp.User.Slug)
	if err != nil {
		logger.Error(
			"Error listing secrets",
			slog.Any("error", err))
		endpoint.InternalServerError(w)
		return
	}

	jsonData, _ := json.Marshal(secret.ListSecretsResponse{
		Secrets: ss,
	})
	_, err = w.Write([]byte(jsonData))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func (cs *SecretService) getSecrets(
	app string,
	req secret.ListSecretsRequest, isAdmin bool, userSlug string,
) ([]*secret.Secret, error) {
	filters := []datastore.Filter{
		datastore.Equals([]string{"app"}, app),
	}

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

		s.Value, err = secretssdk.Decrypt(s.Value, cs.options.SecretEncryptionKey)
		if err != nil {
			return nil, errors.Wrap(err, "failed to decrypt secret")
		}

		secrets = append(secrets, s)
	}

	return secrets, nil
}
