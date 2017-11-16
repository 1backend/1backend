package domain

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (ud *UserDao) GetById(id string) (User, error) {
	u := User{Id: id}
	err := ud.db.Preload("Tokens").First(&u).Error
	return u, err
}

func (ud *UserDao) GetByIds(ids []string) ([]User, error) {
	us := []User{}
	err := ud.db.Where("users.id in (?)", ids).
		Find(&us).Error
	return us, err
}

func (ud *UserDao) GetByEmail(email string) (User, error) {
	u := User{}
	err := ud.db.Where("users.email = ?", email).
		First(&u).Error
	return u, err
}

func (ud *UserDao) GetByNick(nick string) (User, error) {
	u := User{}
	err := ud.db.Where("users.nick = ?", nick).
		First(&u).Error
	return u, err
}

func (ud *UserDao) Create(u User) error {
	return ud.db.Create(&u).Error
}

func (ud *UserDao) Update(u User) error {
	return ud.db.Save(&u).Error
}

// fck this
func (ud *UserDao) OLDUpdate(u User) error {
	oldu, err := ud.GetById(u.Id)
	if err != nil {
		return err
	}
	if u.Name != "" {
		oldu.Name = u.Name
	}
	if u.Password != "" {
		passw, err := bcrypt.GenerateFromPassword([]byte(u.Password), 5)
		if err != nil {
			return err
		}
		oldu.Password = string(passw)
	}
	return ud.db.Save(&oldu).Error
}
