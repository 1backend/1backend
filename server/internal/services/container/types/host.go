package container_svc

type GetHostRequest struct{}

type GetHostResponse struct {
	Host string `json:"host" binding:"required"`
}
