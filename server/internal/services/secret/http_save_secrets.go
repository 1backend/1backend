/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package secretservice

import (
	"context"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"log/slog"
	"net/http"
	"strings"

	sdk "github.com/1backend/1backend/sdk/go"
	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	"github.com/1backend/1backend/sdk/go/secrets"
	secret "github.com/1backend/1backend/server/internal/services/secret/types"
	"github.com/pkg/errors"
	"golang.org/x/crypto/blake2b"
)

// @Id saveSecrets
// @Summary Save Secrets
// @Description Save secrets if authorized to do so
// @Tags Secret Svc
// @Accept json
// @Produce json
// @Param body body secret.SaveSecretsRequest true "Save Secret Request"
// @Success 200 {object} secret.SaveSecretsResponse "Save Secret Response"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /secret-svc/secrets [put]
func (cs *SecretService) SaveSecrets(
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
	if !isAuthRsp.GetAuthorized() {
		endpoint.Unauthorized(w)
		return
	}

	req := &secret.SaveSecretsRequest{}
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		logger.Error(
			"Failed to decode request body",
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

	err = cs.saveSecrets(
		r.Context(),
		req.Secrets,
		isAdmin,
		isAuthRsp.User.Slug,
	)
	if err != nil {
		logger.Error(
			"Failed to save secrets",
			slog.String("user", isAuthRsp.User.Slug),
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	jsonData, _ := json.Marshal(secret.SaveSecretsResponse{})
	_, err = w.Write([]byte(jsonData))
	if err != nil {
		logger.Error("Error writing response", slog.Any("error", err))
		return
	}
}

func (cs *SecretService) saveSecrets(
	ctx context.Context,
	ss []*secret.Secret,
	isAdmin bool,
	userSlug string,
) error {
	cs.options.Lock.Acquire(ctx, "secret-svc-save")
	defer cs.options.Lock.Release(ctx, "secret-svc-save")

	for _, s := range ss {
		if s.App == "" {
			s.App = defaultApp
		}

		secretI, found, err := cs.secretStore.Query(
			datastore.Equals([]string{"app"}, s.App),
			datastore.Equals([]string{"key"}, s.Key),
		).
			FindOne()
		if err != nil {
			return err
		}
		if !found {
			// When secret key does not exist, it can be "claimed" by any authorized user
			// but non admins can only claim keys prefixed with their user slug
			if !isAdmin && !strings.HasPrefix(s.Key, userSlug) {
				return errors.New("users can only claim secrets prefixed with their user slug")
			}
			if s.Id == "" {
				s.Id = sdk.Id("secr")
			}

			// Admin slugs don't need to be saved the writers, readers and deleters blocks
			// as they can read write and anything.
			if !isAdmin {
				if s.Writers == nil {
					s.Writers = []string{userSlug}
				}
				if s.Readers == nil {
					s.Readers = []string{userSlug}
				}
				if s.Deleters == nil {
					s.Deleters = []string{userSlug}
				}
				if s.CanChangeReaders == nil {
					s.CanChangeReaders = []string{userSlug}
				}
				if s.CanChangeWriters == nil {
					s.CanChangeWriters = []string{userSlug}
				}
				if s.CanChangeDeleters == nil {
					s.CanChangeDeleters = []string{userSlug}
				}
			}
			if !s.Encrypted {
				s.Value, err = secrets.Encrypt(s.Value, cs.options.SecretEncryptionKey)
				if err != nil {
					return errors.Wrap(err, "failed to encrypt secret")
				}
			}
			err = cs.checkSum(s)
			if err != nil {
				return errors.Wrap(err, "checksum failed")
			}

			return cs.secretStore.Upsert(s)
		}

		secret := secretI.(*secret.Secret)

		// When a secret is found, it can only be modified by authorized users
		canSave := isAdmin
		if !canSave {
			for _, writer := range secret.Writers {
				if writer == userSlug {
					canSave = true
					break
				}
			}
		}

		if !canSave {
			continue
		}

		if !s.Encrypted {
			s.Value, err = secrets.Encrypt(s.Value, cs.options.SecretEncryptionKey)
			if err != nil {
				return errors.Wrap(err, "failed to encrypt secret")
			}
		}

		err = cs.checkSum(s)
		if err != nil {
			return errors.Wrap(err, "checksum failed")
		}

		err = cs.secretStore.Upsert(s)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cs SecretService) checkSum(s *secret.Secret) error {
	if s.Encrypted && s.Checksum != "" {
		val, err := secrets.Decrypt(s.Value, cs.options.SecretEncryptionKey)
		if err != nil {
			return errors.Wrap(err, "failed to decrypt to compare checksum")
		}

		hash := ""
		switch s.ChecksumAlgorithm {
		case secret.ChecksumAlgorithmSha256:
			h := sha256.Sum256([]byte(val))
			hash = hex.EncodeToString(h[:])

		case secret.ChecksumAlgorithmSha512:
			h := sha512.Sum512([]byte(val))
			hash = hex.EncodeToString(h[:])

		case secret.ChecksumAlgorithmBlake2s:
			h := blake2b.Sum256([]byte(val))
			hash = hex.EncodeToString(h[:])

		case secret.ChecksumAlgorithmUnspecified, secret.ChecksumAlgorithmCRC32:
			fallthrough
		default:
			h := crc32.ChecksumIEEE([]byte(val))
			hash = hex.EncodeToString([]byte{
				byte(h >> 24), byte(h >> 16), byte(h >> 8), byte(h),
			})
		}

		if hash != s.Checksum {
			return fmt.Errorf("checksum incorrect")
		}
	}

	return nil
}
