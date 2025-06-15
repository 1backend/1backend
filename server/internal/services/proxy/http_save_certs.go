/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package proxyservice

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	"github.com/1backend/1backend/sdk/go/secrets"
	"github.com/pkg/errors"

	proxy "github.com/1backend/1backend/server/internal/services/proxy/types"
)

// @ID saveCerts
// @Summary Save Certs
// @Description This endpoint only exist for testing purposes. Only callable by admins
// @Description Certs should be saved by the Proxy Svc and its edge proxying functionality internally,
// @Description not through this endpoint.
// @Tags Proxy Svc
// @Accept json
// @Produce json
// @Param body body proxy.SaveCertsRequest true "Save Certs Request"
// @Success 200 {object} proxy.SaveCertsResponse "Certs saved successfully"
// @Failure 400 {object} proxy.ErrorResponse "Invalid JSON"
// @Failure 401 {object} proxy.ErrorResponse "Unauthorized"
// @Failure 500 {object} proxy.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /proxy-svc/certs [put]
func (cs *ProxyService) SaveCerts(w http.ResponseWriter, r *http.Request) {
	isAuthRsp, statusCode, err := cs.options.PermissionChecker.HasPermission(
		r,
		proxy.PermissionCertEdit,
	)
	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.GetAuthorized() {
		endpoint.Unauthorized(w)
		return
	}

	req := &proxy.SaveCertsRequest{}
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

	_, err = cs.saveCerts(req)
	if err != nil {
		logger.Error(
			"Failed to save certs",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
		return
	}

	endpoint.WriteJSON(w, http.StatusOK, &proxy.SaveCertsResponse{})
}

func (cs *ProxyService) saveCerts(req *proxy.SaveCertsRequest) ([]proxy.Cert, error) {
	if len(req.Certs) == 0 {
		return nil, errors.New("No certs provided")
	}

	rows := make([]datastore.Row, 0, len(req.Certs))
	certs := make([]proxy.Cert, 0, len(req.Certs))

	ids := []any{}
	for _, cert := range req.Certs {
		if cert.Id == "" {
			return nil, errors.New("cert ID is required")
		}
		ids = append(ids, cert.Id)
	}

	existingCerts, err := cs.certStore.Query(
		datastore.IsInList(
			[]string{"id"},
			ids...,
		),
	).Find()
	if err != nil {
		return nil, errors.Wrap(err, "failed to query existing certs")
	}

	existingCertMap := make(map[string]*proxy.Cert, len(existingCerts))
	for _, certI := range existingCerts {
		cert, ok := certI.(*proxy.Cert)
		if !ok {
			return nil, errors.New("invalid cert type in existing certs")
		}
		existingCertMap[cert.Id] = cert
	}

	now := time.Now()

	for _, cert := range req.Certs {
		if cert.Id == "" {
			return nil, errors.New("cert ID is required")
		}
		if cert.Cert == "" {
			return nil, errors.New("cert pem bundle is required")
		}

		r := &proxy.Cert{
			Id:   cert.Id,
			Cert: cert.Cert,
		}

		if existingCert, exists := existingCertMap[cert.Id]; exists {
			r = existingCert
			r.UpdatedAt = now
			r.Cert = cert.Cert

			err = amendCertInfo(r)
			if err != nil {
				logger.Error(
					"Failed to amend cert info",
					slog.String("certId", r.Id),
					slog.Any("error", err),
				)
			}
		} else {
			r.CreatedAt = now
			r.UpdatedAt = now
		}

		r.Cert, err = secrets.Encrypt(r.Cert, cs.options.SecretEncryptionKey)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to encrypt cert '%s'", r.Id)
		}

		rows = append(rows, r)
		certs = append(certs, *r)
	}

	err = cs.certStore.UpsertMany(rows)
	if err != nil {
		return nil, errors.Wrap(err, "failed to save certs")
	}

	return certs, nil
}
