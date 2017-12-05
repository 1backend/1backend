package endpoints

import (
	"github.com/1backend/1backend/backend/state"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

// NewEndpoints is just below the http handlers
func NewEndpoints(
	db *gorm.DB,
	redisClient *redis.Client,
) *Endpoints {
	return &Endpoints{
		db:    db,
		state: state.NewState(redisClient),
	}
}

// Endpoints represents all endpoints of the http server
type Endpoints struct {
	db    *gorm.DB
	state *state.State
}
