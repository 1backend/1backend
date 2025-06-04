package proxy_svc

type Route struct {
	Id        string `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Host   string `json:"host"`
	Target string `json:"target"`
}
