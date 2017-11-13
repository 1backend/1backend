package handlers

import (
	"errors"

	"github.com/crufter/1backend/backend/domain"
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
