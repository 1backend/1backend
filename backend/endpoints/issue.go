package endpoints

import (
	"time"

	"github.com/crufter/1backend/backend/domain"
)

func (e Endpoints) CreateIssue(iss *domain.Issue) error {
	iss.Id = domain.Sid.MustGenerate()
	iss.CreatedAt = time.Now()
	iss.UpdatedAt = time.Now()
	return e.db.Save(iss).Error
}

func (e Endpoints) UpdateIssue(iss *domain.Issue) error {
	iss.UpdatedAt = time.Now()
	return e.db.Save(iss).Error
}

func (e Endpoints) CreateComment(iss *domain.Comment) error {
	iss.Id = domain.Sid.MustGenerate()
	iss.CreatedAt = time.Now()
	iss.UpdatedAt = time.Now()
	return e.db.Save(iss).Error
}

func (e Endpoints) UpdateComment(iss *domain.Comment) error {
	iss.UpdatedAt = time.Now()
	return e.db.Save(iss).Error
}
