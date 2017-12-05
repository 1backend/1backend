package handlers

import (
	"net/http"

	httpr "github.com/julienschmidt/httprouter"
)

func (h *Handlers) PutStar(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	inp := struct {
		Token     string
		ProjectId string
	}{}
	if err := readJsonBody(r, &inp); err != nil {
		write400(w, err)
		return
	}
	user, err := h.ep.GetUser(inp.Token)
	if err != nil {
		write400(w, err)
		return
	}
	err = h.ep.PutStar(user.Id, inp.ProjectId)
	if err != nil {
		write400(w, err)
		return
	}
	write(w, map[string]string{})
}

func (h *Handlers) DeleteStar(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	user, err := h.ep.GetUser(r.URL.Query().Get("token"))
	if err != nil {
		write400(w, err)
		return
	}
	err = h.ep.DeleteStar(user.Id, r.URL.Query().Get("projectId"))
	if err != nil {
		write400(w, err)
		return
	}
	write(w, map[string]string{})
}
