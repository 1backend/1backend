package user_svc

import "time"

// Activation stores the active organization for a specific device.
type Activation struct {
	InternalId string `json:"internalId" swagger:"ignore"`

	AppId string `json:"appId,omitempty"`

	Id string `json:"id,omitempty"`

	CreatedAt time.Time `json:"createdAt" binding:"required"`
	UpdatedAt time.Time `json:"updatedAt" binding:"required"`

	UserId         string `json:"userId,omitempty"`
	Device         string `json:"device,omitempty"`
	OrganizationId string `json:"organizationId,omitempty"`
}

func (o *Activation) GetId() string {
	if o.InternalId == "" {
		panic("organization has no internal id")
	}

	return o.InternalId
}
