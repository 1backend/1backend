package endpoints

import (
	"errors"
	"fmt"
	"time"

	"regexp"

	"github.com/1backend/1backend/backend/domain"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var reg = regexp.MustCompile("^[0-9a-z-]+$")

func (e Endpoints) createToken(tx *gorm.DB, userId string) (*domain.AccessToken, error) {
	token := domain.AccessToken{
		Id:        domain.Sid.MustGenerate(),
		Token:     uuid.Must(uuid.NewV4()).String(),
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
		Token:     uuid.Must(uuid.NewV4()).String(),
		UserId:    user.Id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	defaultServiceToken := &domain.Token{
		Id:          domain.Sid.MustGenerate(),
		Name:        "default",
		Description: "The default token is here to collect quota. We advise you to not use this anywhere.",
		Token:       uuid.Must(uuid.NewV4()).String(),
		UserId:      user.Id,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	testServicetoken := &domain.Token{
		Id:          domain.Sid.MustGenerate(),
		Name:        "test",
		Description: "The test token is here only for demo purposes.",
		Token:       uuid.Must(uuid.NewV4()).String(),
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
		Token:     uuid.Must(uuid.NewV4()).String(),
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

func (e *Endpoints) ChangePassword(user *domain.User, oldPassword, newPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	if err != nil {
		return errors.New("Old password is incorrect")
	}
	npassw, err := bcrypt.GenerateFromPassword([]byte(newPassword), 5)
	if err != nil {
		return err
	}
	user.Password = string(npassw)
	return e.db.Save(user).Error
}

func (e *Endpoints) SendResetEmail(email string) error {
	if email == "" {
		return errors.New("Email can not be empty.")
	}
	user, err := domain.NewUserDao(e.db).GetByEmail(email)
	if err != nil {
		return err
	}
	secret := uuid.Must(uuid.NewV4()).String() // using the uuid v4 here to ensure that it's absolutely unenumerable. org secrets are ok due to email confirmations
	tx := e.db.Begin()
	err = domain.NewResetDao(e.db).Create(domain.Reset{
		Id:        domain.Sid.MustGenerate(),
		Secret:    secret,
		UserId:    user.Id,
		CreatedAt: time.Now(),
	})
	if err != nil {
		tx.Rollback()
		return err
	}
	err = domain.SendPasswordReset(secret, &user)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (e *Endpoints) ResetPassword(secret, newPassword string) (*domain.AccessToken, error) {
	if secret == "" {
		return nil, errors.New("Secret can not be empty.")
	}
	if newPassword == "" {
		return nil, errors.New("New password can not be empty")
	}
	reset, err := domain.NewResetDao(e.db).GetBySecret(secret)
	if err != nil {
		return nil, err
	}
	if reset.CreatedAt.Before(time.Now().Add(-1 * 168 * time.Hour)) {
		return nil, errors.New("Reset code older than a week")
	}
	if reset.Used {
		return nil, errors.New("Reset code already used")
	}
	user, err := domain.NewUserDao(e.db).GetById(reset.UserId)
	if err != nil {
		return nil, err
	}
	passw, err := bcrypt.GenerateFromPassword([]byte(newPassword), 5)
	if err != nil {
		return nil, err
	}
	user.Password = string(passw)
	tx := e.db.Begin()
	err = domain.NewUserDao(tx).Update(user)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	reset.Used = true
	err = domain.NewResetDao(tx).Update(reset)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tokenDao := domain.NewAccessTokenDao(tx)
	token := domain.AccessToken{
		Id:        domain.Sid.MustGenerate(),
		Token:     uuid.Must(uuid.NewV4()).String(),
		UserId:    user.Id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := tokenDao.Create(token); err != nil {
		tx.Rollback()
		return nil, err
	}
	return &token, tx.Commit().Error
}
