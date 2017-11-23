package endpoints

import (
	"errors"
	"fmt"
	"time"

	"regexp"

	"github.com/1backend/1backend/backend/domain"
	"github.com/1backend/1backend/backend/state"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var reg = regexp.MustCompile("^[0-9a-z-]+$")

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

func (e Endpoints) createToken(tx *gorm.DB, userId string) (*domain.AccessToken, error) {
	token := domain.AccessToken{
		Id:        domain.Sid.MustGenerate(),
		Token:     uuid.NewV4().String(),
		UserId:    userId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return &token, domain.NewAccessTokenDao(tx).Create(token)
}

func (e *Endpoints) Register(email, name, password string) (*domain.User, *domain.AccessToken, error) {
	if email == "" {
		return nil, nil, errors.New("Email can not be empty.")
	}
	if password == "" {
		return nil, nil, errors.New("Password can not be empty.")
	}
	if name == "" {
		return nil, nil, errors.New("Name can't be empty.")
	}
	if !reg.Match([]byte(name)) {
		return nil, nil, errors.New("Allowed name characters: lowercase letters, numbers and dash")
	}
	user, err := domain.NewUserDao(e.db).GetByEmail(email)
	if err == nil && user.Id != "" {
		return nil, nil, errors.New("This email is already registered. Try to log in.")
	}
	passw, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	if err != nil {
		return nil, nil, err
	}
	user = domain.User{
		Id:       domain.Sid.MustGenerate(),
		Email:    email,
		Nick:     name,
		Password: string(passw),
	}
	tx := e.db.Begin()
	userDao := domain.NewUserDao(tx)
	if err := userDao.Create(user); err != nil {
		tx.Rollback()
		return nil, nil, fmt.Errorf("error creating new user %s", err.Error())
	}
	tokenDao := domain.NewAccessTokenDao(tx)
	token := domain.AccessToken{
		Id:        domain.Sid.MustGenerate(),
		Token:     uuid.NewV4().String(),
		UserId:    user.Id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	defaultServiceToken := &domain.Token{
		Id:          domain.Sid.MustGenerate(),
		Name:        "default",
		Description: "The default token is here to collect quota. We advise you to not use this anywhere.",
		Token:       uuid.NewV4().String(),
		UserId:      user.Id,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	testServicetoken := &domain.Token{
		Id:          domain.Sid.MustGenerate(),
		Name:        "test",
		Description: "The test token is here only for demo purposes.",
		Token:       uuid.NewV4().String(),
		UserId:      user.Id,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := tx.Create(defaultServiceToken).Error; err != nil {
		tx.Rollback()
		return nil, nil, fmt.Errorf("error creating token for user: %s", err.Error())
	}
	if err := tx.Create(testServicetoken).Error; err != nil {
		tx.Rollback()
		return nil, nil, fmt.Errorf("error creating token for user: %s", err.Error())
	}
	if err := e.state.SetQuota(defaultServiceToken.Token, 95000); err != nil {
		tx.Rollback()
		return nil, nil, fmt.Errorf("error setting quota: %v", err.Error())
	}
	if err := e.state.SetQuota(testServicetoken.Token, 5000); err != nil {
		tx.Rollback()
		return nil, nil, fmt.Errorf("error setting quota: %v", err.Error())
	}
	if err := tokenDao.Create(token); err != nil {
		tx.Rollback()
		return nil, nil, fmt.Errorf("error creating token for user: %s", err.Error())
	}
	err = tx.Commit().Error
	if err != nil {
		return nil, nil, err
	}
	return &user, &token, nil
}

func (e *Endpoints) Login(email, password string) (*domain.User, *domain.AccessToken, error) {
	if email == "" {
		return nil, nil, errors.New("Email can not be empty.")
	}
	if password == "" {
		return nil, nil, errors.New("Password can not be empty.")
	}
	user, err := domain.NewUserDao(e.db).GetByEmail(email)
	if err != nil {
		return nil, nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, nil, errors.New("Could not log in")
	}
	tokenDao := domain.NewAccessTokenDao(e.db)
	token := domain.AccessToken{
		Id:        domain.Sid.MustGenerate(),
		Token:     uuid.NewV4().String(),
		UserId:    user.Id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := tokenDao.Create(token); err != nil {
		return nil, nil, fmt.Errorf("error creating token for user: %s", err.Error())
	}
	return &user, &token, nil
}

// GetUser by token
func (e *Endpoints) GetUser(tokenId string) (*domain.User, error) {
	// first get token
	tokenDao := domain.NewAccessTokenDao(e.db)
	t, err := tokenDao.GetByToken(tokenId)
	if err != nil {
		return nil, fmt.Errorf("token (%s) is associated to no users", tokenId)
	}
	userDao := domain.NewUserDao(e.db)
	u, _ := userDao.GetById(t.UserId)
	return &u, nil
}
