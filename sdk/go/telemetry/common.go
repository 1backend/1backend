/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package telemetry

import (
	"reflect"
	"strings"

	"go.opentelemetry.io/otel/attribute"
)

func errorAttributes(err error) []attribute.KeyValue {
	if err == nil {
		return nil
	}
	return []attribute.KeyValue{
		attribute.String("error.type", errorType(err)),
	}
}

func errorType(err error) string {
	if err == nil {
		return ""
	}

	t := strings.TrimPrefix(reflect.TypeOf(err).String(), "*")
	t = strings.TrimPrefix(t, "github.com/1backend/1backend/")
	t = strings.TrimPrefix(t, "github.com/pkg/errors.")
	t = strings.TrimPrefix(t, "errors.")
	t = strings.TrimPrefix(t, "fmt.")
	if t == "" {
		return "error"
	}
	return t
}

func normalizeBackend(backend string) string {
	if strings.TrimSpace(backend) == "" {
		return "localstore"
	}
	return strings.TrimSpace(strings.ToLower(backend))
}
