package endpoints

import (
	"time"

	"github.com/1backend/1backend/backend/domain"
	uuid "github.com/satori/go.uuid"
)

func (e Endpoints) CreateToken(iss *domain.Token) error {
	iss.Id = uuid.NewV4().String()
	iss.CreatedAt = time.Now()
	iss.UpdatedAt = iss.CreatedAt
	return e.db.Save(iss).Error
}
