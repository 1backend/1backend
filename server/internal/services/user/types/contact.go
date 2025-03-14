/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package user_svc

import (
	"time"
)

type ContactPlatform string

const (
	ContactPlatformEmail    ContactPlatform = "email"
	ContactPlatformPhone    ContactPlatform = "phone"
	ContactPlatformTwitter  ContactPlatform = "twitter"
	ContactPlatformLinkedIn ContactPlatform = "linkedin"
)

type Contact struct {
	// The unique identifier, which can be a URL.
	//
	// Example values: "joe12" (openorch username), "twitter.com/thejoe" (twitter url), "joe@joesdomain.com" (email)
	Id string `json:"id,omitempty" example:"twitter.com/thejoe"`

	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	UserId string `json:"userId,omitempty"`

	// Platform of the contact (e.g., "email", "phone", "twitter")
	Platform string `json:"platform,omitempty" example:"twitter"`

	// Value is the platform local unique identifier.
	// Ie. while the `id` of a Twitter contact is `twitter.com/thejoe`, the value will be only `thejoe`.
	// For email and phones the `id` and the `value` will be the same.
	// This field mostly exists for display purposes.
	//
	// Example values: "joe12" (openorch username), "thejoe" (twitter username), "joe@joesdomain.com" (email)
	Value string `json:"value,omitempty" example:"thejoe"`

	// Whether the contact is verified
	Verified bool `json:"verified,omitempty"`

	// If this is the primary contact method
	IsPrimary bool `json:"isPrimary,omitempty"`
}
