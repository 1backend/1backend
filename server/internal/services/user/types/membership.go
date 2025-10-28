package user_svc

import "time"

// Membership links a user to an organization.
// Note: Roles come from Enrollments, not Memberships.
// When created through SaveMembership, the user also receives the
// corresponding Enrollment, enabling their dynamic role assignment
// (e.g. `user-svc:org:{org_123}:user`).
type Membership struct {
	InternalId string `json:"internalId" swaggerignore:"true"`

	AppId string `json:"appId,omitempty"`

	Id string `json:"id,omitempty"`

	CreatedAt time.Time  `json:"createdAt" binding:"required"`
	UpdatedAt time.Time  `json:"updatedAt" binding:"required"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	OrganizationId string `json:"organizationId,omitempty"`
	UserId         string `json:"userId,omitempty"`

	// Active/default organization for a user.
	// There can only be one per user.
	Active bool `json:"active,omitempty"`
}

func (o *Membership) GetId() string {
	return o.Id
}

type SaveMembershipRequest struct {
}

type SaveMembershipResponse struct {
}

type DeleteMembershipRequest struct {
}

type DeleteMembershipResponse struct {
}
