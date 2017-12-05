package handlers

import (
	"errors"

	"net/http"

	"github.com/1backend/1backend/backend/domain"
	httpr "github.com/julienschmidt/httprouter"
)

func (h *Handlers) getUser(tokenId string) (*domain.User, error) {
	token, err := domain.NewAccessTokenDao(h.db).GetByToken(tokenId)
	if err != nil {
		return nil, err
	}
	u, err := domain.NewUserDao(h.db).GetById(token.UserId)
	return &u, err
}

func (h *Handlers) ownsProject(tokenId, projectId string) error {
	token, err := domain.NewAccessTokenDao(h.db).GetByToken(tokenId)
	if err != nil {
		return err
	}
	user, err := domain.NewUserDao(h.db).GetById(token.UserId)
	if err != nil {
		return err
	}
	proj := domain.Project{}
	err = h.db.Select("author").Where("id = ?", projectId).First(&proj).Error
	if err != nil {
		return err
	}
	if proj.Author != user.Nick {
		return errors.New("No right to project")
	}
	return nil
}

func (h *Handlers) hasNick(tokenId, author string) error {
	token, err := domain.NewAccessTokenDao(h.db).GetByToken(tokenId)
	if err != nil {
		return err
	}
	user, err := domain.NewUserDao(h.db).GetById(token.UserId)
	if err != nil {
		return err
	}
	if user.Nick != author {
		return errors.New("No right to access")
	}
	return nil
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	eb := struct {
		Email    string
		Password string
	}{}
	if err := readJsonBody(r, &eb); err != nil {
		write400(w, err)
		return
	}
	user, token, err := h.ep.Login(eb.Email, eb.Password)
	if err != nil {
		write400(w, err)
		return
	}
	write(w, map[string]interface{}{
		"user":  user,
		"token": token,
	})
}

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	eb := struct {
		Email    string
		Password string
		Nick     string
	}{}
	if err := readJsonBody(r, &eb); err != nil {
		write400(w, err)
		return
	}
	user, token, err := h.ep.Register(eb.Email, eb.Nick, eb.Password)
	if err != nil {
		write400(w, err)
		return
	}
	write(w, map[string]interface{}{
		"user":  user,
		"token": token,
	})
}
