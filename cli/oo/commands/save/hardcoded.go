/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package save

var secretMeta = Meta{
	Transport: &Transport{
		Endpoint:    "/secret-svc/secrets",
		Method:      "PUT",
		Body:        "secrets",
		Array:       true,
		ContentType: "application/json",
	},
}

var permitMeta = Meta{
	Transport: &Transport{
		Endpoint:    "/user-svc/permits",
		Method:      "PUT",
		Body:        "permit",
		Array:       true,
		ContentType: "application/json",
	},
}

var routeMeta = Meta{
	Transport: &Transport{
		Endpoint:    "/proxy-svc/routes",
		Method:      "PUT",
		Body:        "route",
		Array:       true,
		ContentType: "application/json",
	},
}
