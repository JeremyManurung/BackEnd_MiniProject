package komentar

import (
	"gorm.io/gorm"
)

type repository struct{
	db *gorm.DB
}

type Repository interface {
	Save(komentar Komentar) (Komentar, error)
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