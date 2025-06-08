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

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"
	"github.com/pkg/errors"

	proxy "github.com/1backend/1backend/server/internal/services/proxy/types"
)

// @ID listCerts
// @Summary List Certs
// @Description List certs that the edge proxy will use to cert requests.
// @Tags Proxy Svc
// @Accept json
// @Produce json
// @Param body body proxy.ListCertsRequest false "List Certs Request"
// @Success 200 {object} proxy.ListCertsResponse "Certs listed successfully"
// @Failure 400 {object} proxy.ErrorResponse "Invalid JSON"
// @Failure 401 {object} proxy.ErrorResponse "Unauthorized"
// @Failure 500 {object} proxy.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /proxy-svc/certs [post]
func (cs *ProxyService) ListCerts(w http.ResponseWriter, r *http.Request) {
	isAuthRsp, statusCode, err := cs.options.PermissionChecker.HasPermission(
		r,
		proxy.PermissionCertView,
	)

	if err != nil {
		endpoint.WriteErr(w, statusCode, err)
		return
	}
	if !isAuthRsp.GetAuthorized() {
		endpoint.Unauthorized(w)
		return
	}

	req := &proxy.ListCertsRequest{}

	if r.ContentLength > 0 {
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
	}

	certs, err := cs.listCerts(req)
	if err != nil {
		logger.Error(
			"Failed to list certs",
			slog.Any("error", err),
		)
		endpoint.InternalServerError(w)
	}

	endpoint.WriteJSON(w, http.StatusOK, &proxy.ListCertsResponse{
		Certs: certs,
	})
}

func (cs *ProxyService) listCerts(req *proxy.ListCertsRequest) ([]proxy.Cert, error) {
	filters := []datastore.Filter{}

	if len(req.Ids) > 0 {
		ids := []any{}
		for _, id := range req.Ids {
			ids = append(ids, id)
		}
		filters = append(filters, datastore.IsInList([]string{"id"}, ids...))
	}

	certIs, err := cs.certStore.Query(filters...).
		Find()
	if err != nil {
		return nil, errors.Wrap(err, "failed to query certs")
	}

	certs := make([]proxy.Cert, 0, len(certIs))
	for _, certI := range certIs {
		cert, ok := certI.(*proxy.Cert)
		if !ok {
			return nil, errors.Errorf("expected cert type, got %T", certI)
		}

		if cert.CommonName == "" {
			err = amendCertInfo(cert)
			if err != nil {
				// Only log, do not error.
				logger.Error(
					"Failed to amend cert info",
					slog.String("certId", cert.Id),
					slog.Any("error", err),
				)
			}
		}

		certs = append(certs, *cert)
	}

	return certs, nil
}

func amendCertInfo(cert *proxy.Cert) error {
	info, err := parseCertInfo([]byte(cert.Cert))
	if err != nil {
		return errors.Wrapf(err, "failed to parse cert info for key '%s'", cert.Id)
	}

	cert.CommonName = info.CommonName
	cert.Issuer = info.Issuer
	cert.NotBefore = info.NotBefore
	cert.NotAfter = info.NotAfter
	cert.SerialNumber = info.SerialNumber
	cert.DNSNames = info.DNSNames
	cert.SerialNumber = info.SerialNumber
	cert.SignatureAlgorithm = info.SignatureAlgorithm
	cert.PublicKeyAlgorithm = info.PublicKeyAlgorithm
	cert.PublicKeyBitLength = info.PublicKeyBitLength

	return nil
}
