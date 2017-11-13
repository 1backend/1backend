package endpoints

import (
	"errors"

	"github.com/1backend/1backend/backend/domain"
)

//
// Utility functions for auth, not exposed as endpoints (yet)
//

func (e *Endpoints) OwnsProject(tokenId, projectId string) error {
	token, err := domain.NewAccessTokenDao(e.db).GetByToken(tokenId)
	if err != nil {
		return err
	}
	user, err := domain.NewUserDao(e.db).GetById(token.UserId)
	if err != nil {
		return err
	}
	proj := domain.Project{}
	err = e.db.Select("author").Where("id = ?", projectId).First(&proj).Error
	if err != nil {
		return err
	}
	if proj.Author != user.Nick {
		return errors.New("No right to project")
	}
	return nil
}

func (e *Endpoints) HasNick(tokenId, author string) error {
	token, err := domain.NewAccessTokenDao(e.db).GetByToken(tokenId)
	if err != nil {
		return err
	}
	user, err := domain.NewUserDao(e.db).GetById(token.UserId)
	if err != nil {
		return err
	}
	if user.Nick != author {
		return errors.New("No right to access")
	}
	return nil
}
