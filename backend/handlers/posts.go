package handlers

import (
	"errors"
	"net/http"

	httpr "github.com/julienschmidt/httprouter"

	"github.com/1backend/1backend/backend/domain"
)

func (h *Handlers) GetPosts(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	nick := r.URL.Query().Get("nick")
	user, err := domain.NewUserDao(h.db).GetByNick(nick)
	if err != nil {
		write400(w, err)
		return
	}
	posts := []domain.Post{}
	err = h.db.Where("user_id = ?", user.Id).Preload("User").Find(&posts).Error
	if err != nil {
		write400(w, err)
		return
	}
	write(w, posts)
}

func (h *Handlers) GetPost(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	postId := r.URL.Query().Get("postId")
	post := domain.Post{}
	err := h.db.Where("id = ?", postId).Preload("User").Find(&post).Error
	if err != nil {
		write400(w, err)
		return
	}
	write(w, post)
}

func (h *Handlers) CreatePost(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	inp := struct {
		Token string
		Post  domain.Post
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
	inp.Post.UserId = user.Id
	err = h.ep.CreatePost(&inp.Post)
	if err != nil {
		write400(w, err)
		return
	}
	write(w, map[string]string{})
}

func (h *Handlers) UpdatePost(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	inp := struct {
		Token string
		Post  domain.Post
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
	iss := domain.Post{}
	err = h.db.Where("id = ?", inp.Post.Id).First(&iss).Error
	if err != nil {
		write400(w, err)
		return
	}
	if user.Id != iss.UserId {
		write400(w, errors.New("No right"))
		return
	}
	err = h.ep.UpdatePost(&inp.Post)
	if err != nil {
		write500(w, err)
		return
	}
	write(w, map[string]string{})
}
