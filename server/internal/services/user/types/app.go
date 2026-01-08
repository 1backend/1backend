package user_svc

// App identifies a tenant namespace in 1Backend.
//
// The Host field is keyed primarily by hostname (e.g. "example.com").
// For local or non-standard deployments, port numbers are also taken
// into account (e.g. "example.com:8080", "[::1]:11337").
//
// In practice, public edge requests will almost always be bare hostnames
// because standard ports (80/443) are implicit in the Host header.
type App struct {
	Id   string `json:"id" binding:"required"`
	Host string `json:"host" binding:"required"`
}

func (a *App) GetId() string {
	return a.Id
}

type ListAppsRequest struct {
	Ids   []string `json:"ids,omitempty"`
	Hosts []string `json:"host,omitempty"`
}

type ListAppsResponse struct {
	Apps []App `json:"apps" binding:"required"`
}

type ReadAppRequest struct {
	Host string `json:"host,omitempty"`
}

type ReadAppResponse struct {
	App App `json:"app" binding:"required"`
}

type UpdateAppRequest struct {
	Id      string `json:"id" binding:"required"`
	NewHost string `json:"newHost" binding:"required"`
}

type UpdateAppResponse struct {
}
