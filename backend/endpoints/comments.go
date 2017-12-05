package endpoints

import (
	"time"

	"github.com/1backend/1backend/backend/domain"
)

func (e Endpoints) CreateComment(iss *domain.Comment) error {
	iss.Id = domain.Sid.MustGenerate()
	iss.CreatedAt = time.Now()
	iss.UpdatedAt = iss.CreatedAt
	return e.db.Save(iss).Error
}

func (e Endpoints) UpdateComment(iss *domain.Comment) error {
	iss.UpdatedAt = time.Now()
	return e.db.Save(iss).Error
}
