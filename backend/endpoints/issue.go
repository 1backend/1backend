package endpoints

import (
	"time"

	"github.com/1backend/1backend/backend/domain"
	uuid "github.com/satori/go.uuid"
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

func (e Endpoints) CreateToken(iss *domain.Token) error {
	iss.Id = uuid.NewV4().String()
	iss.CreatedAt = time.Now()
	iss.UpdatedAt = time.Now()
	return e.db.Save(iss).Error
}
