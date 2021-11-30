package transaksi

import (
	"minipro/bantuan"
)

type service struct {
	repository         Repository
	bantuanRepository  bantuan.Repository
}

type Service interface {
	GetTransaksisByBantuanID(input int) ([]Transaksi, error)
	GetTransaksisByUserID(userID int) ([]Transaksi, error)
}

func NewService(repository Repository, bantuanRepository bantuan.Repository) *service {
	return &service{repository, bantuanRepository}
}

func (s *service) GetTransaksisByBantuanID(input int) ([]Transaksi, error) {
	
	// bantuan, err := s.bantuanRepository.FindByID(input)
	// if err != nil{
	// 	return []Transaksi{}, err
	// }

	// if bantuan.UserID != input.User.ID

	transaksis, err := s.repository.GetByBantuanID(input)
	if err != nil {
		return []Transaksi{}, err
	}

	return transaksis, nil
}

func (s *service) GetTransaksisByUserID(userID int) ([]Transaksi, error){
	transaksis, err := s.repository.GetByUserID(userID)
	if err != nil{
		return transaksis, err
	}

	return transaksis, nil

}
