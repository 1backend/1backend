package handlers

import (
	"errors"
	"net/http"

	"github.com/1backend/1backend/backend/domain"
	httpr "github.com/julienschmidt/httprouter"
)

func (h *Handlers) CreateComment(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	inp := struct {
		Token   string
		Comment domain.Comment
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
	inp.Comment.UserId = user.Id
	err = h.ep.CreateComment(&inp.Comment)
	if err != nil {
		write400(w, err)
		return
	}
	write(w, map[string]string{})
}

func (h *Handlers) UpdateComment(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	inp := struct {
		Token   string
		Comment domain.Comment
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
	comm := domain.Comment{}
	err = h.db.Where("id = ?", inp.Comment.Id).First(&comm).Error
	if err != nil {
		write400(w, err)
		return
	}
	if user.Id != comm.UserId {
		write400(w, errors.New("No right"))
		return
	}
	err = h.ep.UpdateComment(&inp.Comment)
	if err != nil {
		write500(w, err)
		return
	}
	write(w, map[string]string{})
}
