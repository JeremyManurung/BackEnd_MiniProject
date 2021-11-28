package bantuan

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Bantuan, error)
	FindByUserID(userID int) ([]Bantuan, error)
	FindByID(ID int) (Bantuan, error)
	Save(bantuan Bantuan) (Bantuan, error)
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

func (r *repository) FindByID(ID int) (Bantuan, error) {
	var bantuan Bantuan

	err := r.db.Preload("User").Preload("BantuanImgs").Where("id = ?", ID).Find(&bantuan).Error

	if err != nil {
		return bantuan, err
	}

	return bantuan, nil
}

func (r *repository) Save(bantuan Bantuan) (Bantuan, error) {
	err := r.db.Create(&bantuan).Error
	if err != nil {
		return bantuan, err
	}

	return bantuan, nil
}