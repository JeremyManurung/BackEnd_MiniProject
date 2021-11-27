package bantuan

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Bantuan, error)
	FindByUserID(userID int) ([]Bantuan, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Bantuan, error) {
	var bantuans []Bantuan

	err := r.db.Preload("BantuanImgs", "bantuan_imgs.img_utama = 1").Find(&bantuans).Error
	if err != nil {
		return bantuans, err
	}

	return bantuans, nil
}

func (r *repository) FindByUserID(userID int) ([]Bantuan, error) {
	var bantuans []Bantuan

	err := r.db.Where("user_id = ?", userID).Preload("BantuanImgs", "bantuan_imgs.img_utama = 1").Find(&bantuans).Error
	if err != nil {
		return bantuans, err
	}

	return bantuans, nil
}