package redirect

type RedirectInput struct {
	Id         string `json:"id" yaml:"id"`
	Target     string `json:"target" yaml:"target"`
	StatusCode int    `json:"statusCode,omitempty" yaml:"statusCode,omitempty"`
}

type SaveRedirectsRequest struct {
	Redirects []RedirectInput `json:"redirects"`
}

type ListRedirectsRequest struct {
	Ids []string `json:"ids,omitempty"`
}

type ListRedirectsResponse struct {
	Redirects []RedirectInput `json:"redirects"`
}
