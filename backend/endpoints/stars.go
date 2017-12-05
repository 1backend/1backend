package endpoints

import (
	"github.com/1backend/1backend/backend/domain"
)

func (e Endpoints) PutStar(userId, projectId string) error {
	star := &domain.Star{
		Id:        domain.Sid.MustGenerate(),
		ProjectId: projectId,
		UserId:    userId,
	}
	tx := e.db.Begin()
	err := tx.Save(star).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Exec("update projects set stars = stars + 1 where id = ?", projectId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (e Endpoints) DeleteStar(userId, projectId string) error {
	star := domain.Star{}
	err := e.db.Where("user_id = ? AND project_id = ?", userId, projectId).First(&star).Error
	if err != nil {
		return err
	}
	tx := e.db.Begin()
	err = tx.Delete(star).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Exec("update projects set stars = stars - 1 where id = ?", projectId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
