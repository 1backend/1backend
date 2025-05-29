/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package proxyservice

import (
	"context"
	"io"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/logger"
)

func (cs *ProxyService) Route(w http.ResponseWriter, r *http.Request) {
	statusCode, err := cs.route(w, r)

	if err != nil {
		logger.Error("Error proxying OPTIONS request",
			slog.String("path", r.URL.Path),
			slog.String("error", err.Error()),
		)
		if r.Method == http.MethodOptions {
			// We don't want to write errors to response on OPTIONS requests
			// because the response won't be visible in the chrome dev tools.
			// We log it instead.

			return
		}

		w.WriteHeader(statusCode)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			logger.Error("Error writing response",
				slog.String("error", err.Error()),
			)
		}
		return
	}
}

func (cs *ProxyService) route(w http.ResponseWriter, r *http.Request) (int, error) {
	logger.Debug("Proxying",
		slog.String("path", r.URL.Path),
		slog.String("method", r.Method),
	)

	// The proxy service's LazyStart() must be called here because OPTIONS requests
	// are not handled the standard Lazy logic. Unlike other services, the Proxy does handle
	// OPTIONS requests and requires initialization (including token acquisition) to do so.
	err := cs.LazyStart()
	if err != nil && r.Method != http.MethodOptions {
		return http.StatusInternalServerError, errors.Wrap(err, "error starting proxy service")
	}

	// @todo cache?

	serviceSlug := getServiceSlug(r)

	rsp, _, err := cs.options.ClientFactory.Client(client.WithToken(cs.token)).
		RegistrySvcAPI.ListInstances(context.Background()).
		Slug(serviceSlug).
		Execute()
	if err != nil {
		return http.StatusInternalServerError, errors.Wrap(err, "error listing instances")
	}

	if len(rsp.Instances) == 0 {
		return http.StatusNotFound, errors.New("no instances found")
	}

	// Prioritize healthy instances
	selectedInstances := make([]openapi.RegistrySvcInstance, 0, len(rsp.Instances))
	for _, instance := range rsp.Instances {
		if instance.Status == openapi.InstanceStatusHealthy {
			selectedInstances = append(selectedInstances, instance)
		}
	}

	// But fall back to any instances if none found
	if len(selectedInstances) == 0 {
		selectedInstances = rsp.Instances
	}

	randomIndex := rand.IntN(len(selectedInstances))
	instance := selectedInstances[randomIndex]

	uri := strings.TrimSuffix(instance.Url, "/") +
		"/" +
		strings.TrimLeft(r.URL.Path, "/")
	if r.URL.RawQuery != "" {
		uri += "?" + r.URL.RawQuery
	}

	req, err := http.NewRequest(r.Method, uri, r.Body)
	if err != nil {
		return http.StatusInternalServerError, errors.Wrap(err, "error creating proxy request")
	}

	req.Header = r.Header

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return http.StatusInternalServerError, errors.Wrap(err, "error proxying request")
	}

	defer resp.Body.Close()

	for k, v := range resp.Header {
		// Skip "Content-Length" and "Transfer-Encoding" to prevent HTTP response errors.
		// - Go automatically sets "Content-Length" based on the actual body size.
		// - "Transfer-Encoding: chunked" conflicts with "Content-Length", so we avoid copying it.
		if k == "Content-Length" || k == "Transfer-Encoding" {
			continue
		}
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}

	logger.Debug("Proxy request returned",
		slog.Int("statusCode", resp.StatusCode),
	)

	w.WriteHeader(resp.StatusCode)

	// Unfortunately io.Copy does not work here and causes "invalid write result"

	//_, err = io.Copy(w, resp.Body)
	//if err != nil {
	//	logger.Error("Error proxying body",
	//		slog.String("error", err.Error()),
	//	)
	//}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return http.StatusInternalServerError, errors.Wrap(err, "failed to read body while proxying")
	}

	_, err = w.Write(body)
	if err != nil {
		logger.Error("Error writing response body", slog.String("error", err.Error()))
	}

	if flusher, ok := w.(http.Flusher); ok {
		flusher.Flush()
	}

	return http.StatusOK, nil
}

// gets service slug from http request path
// eg. /my-svc/my-endpoint -> my-svc
func getServiceSlug(r *http.Request) string {
	cleanedPath := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(cleanedPath, "/")
	if len(parts) > 0 && parts[0] != "" {
		return parts[0]
	}
	return ""
}
