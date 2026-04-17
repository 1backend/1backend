package user_svc

import "time"

// Membership links a user to an organization.
// Note: Roles come from Enrollments, not Memberships.
// When created through SaveMembership, the user also receives the
// corresponding Enrollment, enabling their dynamic role assignment
// (e.g. `user-svc:org:{org_123}:user`).
type Membership struct {
	InternalId string `json:"internalId" swagger:"ignore"`

	AppId string `json:"appId,omitempty"`

	Id string `json:"id,omitempty"`

	CreatedAt time.Time  `json:"createdAt" binding:"required"`
	UpdatedAt time.Time  `json:"updatedAt" binding:"required"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	OrganizationId string `json:"organizationId,omitempty"`
	UserId         string `json:"userId,omitempty"`
	Device         string `json:"device,omitempty"`

	// Active/default organization for a user on a specific device.
	// There can only be one active membership per (user, device).
	Active bool `json:"active,omitempty"`
}

func (o *Membership) GetId() string {
	return o.Id
}

type SaveMembershipRequest struct {
	// Device scope of the membership activation. Defaults to `unknown`.
	Device string `json:"device,omitempty"`

	// If true, this membership becomes the active organization for the
	// specified device. Only one membership can remain active per (user, device).
	Active bool `json:"active,omitempty"`
}

type SaveMembershipResponse struct {
}

type DeleteMembershipRequest struct {
}

type DeleteMembershipResponse struct {
}
