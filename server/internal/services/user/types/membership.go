package user_svc

import "time"

// Membership links a user to an organization.
// Membership is the canonical organization relationship for a user.
type Membership struct {
	InternalId string `json:"internalId" swagger:"ignore"`

	AppId string `json:"appId,omitempty"`

	Id string `json:"id,omitempty"`

	CreatedAt time.Time  `json:"createdAt" binding:"required"`
	UpdatedAt time.Time  `json:"updatedAt" binding:"required"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`

	OrganizationId string `json:"organizationId,omitempty"`
	UserId         string `json:"userId,omitempty"`

	Status MembershipStatus `json:"status,omitempty"`

	Roles []string `json:"roles,omitempty"`

	InvitedBy string `json:"invitedBy,omitempty"`

	AcceptedAt *time.Time `json:"acceptedAt,omitempty"`
}

func (o *Membership) GetId() string {
	if o.InternalId == "" {
		panic("membership has no internal id")
	}

	return o.InternalId
}

type MembershipStatus string

const (
	MembershipStatusPending  MembershipStatus = "pending"
	MembershipStatusAccepted MembershipStatus = "accepted"
	MembershipStatusDeclined MembershipStatus = "declined"
)

type SaveMembershipRequest struct {
	Roles []string `json:"roles,omitempty"`
}

type SaveMembershipResponse struct {
	Membership Membership `json:"membership" binding:"required"`
}

type DeleteMembershipRequest struct {
}

type DeleteMembershipResponse struct {
}

type AcceptMembershipRequest struct {
	Activate bool `json:"activate,omitempty"`
}

type AcceptMembershipResponse struct {
	Membership Membership `json:"membership" binding:"required"`
	Token      *Token     `json:"token,omitempty"`
}

type DeclineMembershipRequest struct {
}

type DeclineMembershipResponse struct {
	Membership Membership `json:"membership" binding:"required"`
}

type ListMembershipsRequest struct {
	OrganizationId string           `json:"organizationId,omitempty"`
	UserId         string           `json:"userId,omitempty"`
	Status         MembershipStatus `json:"status,omitempty"`
	Limit          int              `json:"limit,omitempty"`
	AfterTime      time.Time        `json:"afterTime,omitempty"`
}

type ListMembershipsResponse struct {
	Memberships []struct {
		Membership   Membership   `json:"membership" binding:"required"`
		Organization Organization `json:"organization" binding:"required"`
		User         User         `json:"user" binding:"required"`
	} `json:"memberships" binding:"required"`
}
