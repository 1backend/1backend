package domain

import "github.com/jinzhu/gorm"

type ResetDao struct {
	db *gorm.DB
}

func NewResetDao(db *gorm.DB) *ResetDao {
	return &ResetDao{db: db}
}

func (ud *ResetDao) GetById(id string) (Reset, error) {
	u := Reset{Id: id}
	err := ud.db.First(&u).Error
	return u, err
}

func (ud *ResetDao) GetByIds(ids []string) ([]Reset, error) {
	us := []Reset{}
	err := ud.db.Where("resets.id in (?)", ids).
		Find(&us).Error
	return us, err
}

func (ud *ResetDao) GetBySecret(secret string) (Reset, error) {
	u := Reset{}
	err := ud.db.Where("resets.secret = ?", secret).
		First(&u).Error
	return u, err
}

func (ud *ResetDao) Create(u Reset) error {
	return ud.db.Create(&u).Error
}

func (ud *ResetDao) Update(u Reset) error {
	return ud.db.Save(&u).Error
}
