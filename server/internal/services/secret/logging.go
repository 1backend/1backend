/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package secretservice

import (
	"log/slog"
	"net/http"
	"strings"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/auth"
)

func secretRequestLogArgs(
	r *http.Request,
	authRsp *openapi.UserSvcHasPermissionResponse,
	attrs ...slog.Attr,
) []any {
	args := []any{}

	if r != nil {
		args = append(args,
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.String("remoteAddr", r.RemoteAddr),
			slog.String("userAgent", r.UserAgent()),
		)

		token, _ := auth.AuthorizerImpl{}.TokenFromRequest(r)
		for _, attr := range auth.TokenDebugAttrs(token) {
			args = append(args, attr)
		}
	}

	if authRsp != nil {
		args = append(args,
			slog.String("callerSlug", authRsp.User.Slug),
			slog.String("callerUserId", authRsp.User.Id),
			slog.String("callerAppId", authRsp.AppId),
			slog.String("callerAppHost", authRsp.App.Host),
		)
	}

	for _, attr := range attrs {
		args = append(args, attr)
	}

	return args
}

func isRequestAuthError(err error) bool {
	if err == nil {
		return false
	}

	msg := err.Error()
	return strings.Contains(msg, "token is expired") ||
		strings.Contains(msg, "no token found") ||
		strings.Contains(msg, "token is empty") ||
		strings.Contains(msg, "failed to parse JWT") ||
		strings.Contains(msg, "invalid JWT token")
}
