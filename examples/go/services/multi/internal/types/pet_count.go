package multi_svc

type CountPetsRequest struct{}

type CountPetsResponse struct {
	PetCount int `json:"petCount" binding:"required"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
