package transaksi

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetByBantuanID(bantuanID int) ([]Transaksi, error)
	GetByUserID(userID int)([]Transaksi, error)
	Save(transaksi Transaksi) (Transaksi, error)
	Update(transaksi Transaksi) (Transaksi, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByBantuanID(bantuanID int) ([]Transaksi, error) {
	var transaksis []Transaksi

	err := r.db.Preload("User").Where("bantuan_id = ?", bantuanID).Order("id desc").Find(&transaksis).Error
	if err != nil {
		return transaksis, err
	}

	return transaksis, nil
}

func (r *repository) GetByUserID(userID int) ([]Transaksi, error) {
	var transaksis []Transaksi

	err := r.db.Preload("Bantuan.BantuanImgs", "bantuan_imgs.img_utama = 1").Where("user_id = ?", userID).Order("id desc").Find(&transaksis).Error
	if err != nil {
		return transaksis, err
	}

	return transaksis, nil
}

func (r *repository) Save(transaksi Transaksi) (Transaksi, error) {
	err := r.db.Create(&transaksi).Error

	if err != nil {
		return transaksi, err
	}

	return transaksi, nil
	
}

func (r *repository) Update(transaksi Transaksi) (Transaksi, error) {
	err := r.db.Save(&transaksi).Error

	if err != nil {
		return transaksi, err
	}

	return transaksi, nil
	
}