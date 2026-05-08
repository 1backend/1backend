/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package userservice

import (
	"time"

	user "github.com/1backend/1backend/server/internal/services/user/types"
)

func userLifecycleEventPayload(
	app *user.App,
	token *user.Token,
	slug string,
	eventTime time.Time,
) map[string]any {
	if eventTime.IsZero() {
		eventTime = time.Now().UTC()
	}

	payload := map[string]any{
		"appId":  token.AppId,
		"userId": token.UserId,
		"device": token.Device,
		"time":   eventTime.UTC(),
	}

	if slug != "" {
		payload["slug"] = slug
	}
	if app != nil && app.Host != "" {
		payload["appHost"] = app.Host
	} else if token.App != nil && token.App.Host != "" {
		payload["appHost"] = token.App.Host
	}

	return payload
}
