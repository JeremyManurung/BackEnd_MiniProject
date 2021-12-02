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
}

func NewService(repository Repository, bantuanRepository bantuan.Repository) *service {
	return &service{repository, bantuanRepository}
}

func (s *service) GetTransaksisByBantuanID(input int) ([]Transaksi, error) {
	


	transaksis, err := s.repository.GetByBantuanID(input)
	if err != nil {
		return []Transaksi{}, err
	}

	return transaksis, nil
}