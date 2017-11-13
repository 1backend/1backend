package domain

import "github.com/jinzhu/gorm"

type AccessTokenDao struct {
	db *gorm.DB
}

func NewAccessTokenDao(db *gorm.DB) *AccessTokenDao {
	return &AccessTokenDao{db: db}
}

func (at *AccessTokenDao) Create(model AccessToken) error {
	return at.db.Create(&model).Error
}

func (at *AccessTokenDao) GetByToken(token string) (AccessToken, error) {
	model := AccessToken{}
	err := at.db.Where("access_tokens.token = ?", token).
		First(&model).Error
	return model, err
}

func (at *AccessTokenDao) Delete(id string) error {
	return at.db.Delete(&AccessToken{Id: id}).Error
}

func (at *AccessTokenDao) DeleteByToken(token string) error {
	return at.db.Where("access_tokens.token = ?", token).
		Delete(&AccessToken{}).Error
}
