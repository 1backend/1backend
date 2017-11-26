package endpoints

import (
	"errors"
	"time"

	"github.com/1backend/1backend/backend/deploy"
	"github.com/1backend/1backend/backend/domain"
	tp "github.com/1backend/1backend/backend/tech-pack"
)

func (e Endpoints) CreateProject(proj *domain.Project) error {
	if !reg.Match([]byte(proj.Name)) {
		return errors.New("Allowed project name characters: lowercase letters, numbers, dash")
	}
	proj.Description = "Change this to something sensible"
	proj.Id = domain.Sid.MustGenerate()
	proj.InfraPassword = domain.Sid.MustGenerate()
	proj.CreatedAt = time.Now()
	proj.UpdatedAt = time.Now()
	tp.GetProvider(proj).CreateProjectPlugin()
	err := e.db.Save(proj).Error
	if err != nil {
		return err
	}
	for _, v := range proj.Endpoints {
		v.Id = domain.Sid.MustGenerate()
		v.ProjectId = proj.Id
		v.CreatedAt = time.Now()
		v.UpdatedAt = time.Now()
		err = e.db.Save(&v).Error
		if err != nil {
			return err
		}
	}
	for _, v := range proj.Dependencies {
		v.Id = domain.Sid.MustGenerate()
		v.ProjectId = proj.Id
		v.CreatedAt = time.Now()
		v.UpdatedAt = time.Now()
		err = e.db.Save(&v).Error
		if err != nil {
			return err
		}
	}
	go deploy.NewDeployer(e.db).Deploy(proj)
	return nil
}

func (e Endpoints) UpdateProject(proj *domain.Project) error {
	if !reg.Match([]byte(proj.Name)) {
		return errors.New("Allowed project name characters: lowercase letters, numbers, dash")
	}
	proj.InfraPassword = domain.Sid.MustGenerate()
	err := e.db.Save(proj).Error
	if err != nil {
		return err
	}
	for _, v := range proj.Endpoints {
		if v.Id == "" {
			v.Id = domain.Sid.MustGenerate()
		}
		v.ProjectId = proj.Id
		err = e.db.Save(&v).Error
		if err != nil {
			return err
		}
	}
	for _, v := range proj.Dependencies {
		if v.Id == "" {
			v.Id = domain.Sid.MustGenerate()
		}
		v.ProjectId = proj.Id
		err = e.db.Save(&v).Error
		if err != nil {
			return err
		}
	}
	go deploy.NewDeployer(e.db, e.state).Deploy(proj)
	return nil
}

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
