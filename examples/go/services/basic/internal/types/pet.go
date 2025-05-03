package basic_svc

import "time"

type Pet struct {
	Id        string    `json:"id" binding:"required"`
	CreatedAt time.Time `json:"createdAt" binding:"required"`
	UpdatedAt time.Time `json:"updatedAt" binding:"required"`

	Name string `json:"name,omitempty"`
}

func (a Pet) GetId() string {
	return a.Id
}

type SavePetRequest struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name" binding:"required"`
}

type SavePetResponse struct{}

type ListPetsRequest struct {
}

type ListPetsResponse struct {
	Pets []Pet `json:"pets,omitempty"`
}
