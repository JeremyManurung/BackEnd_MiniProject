package komentar

import (
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	Save(komentar Komentar) (Komentar, error)
	Delete(ID int) error
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(komentar Komentar) (Komentar, error) {
	err := r.db.Create(&komentar).Error

	if err != nil {
		return komentar, err
	}

	return komentar, nil

}

func (r *repository) Delete(ID int) error {
	return r.db.Delete(Komentar{}, ID).Error
}
