package transaksi

import (
	"minipro/bantuan"
	"minipro/pembayaran"
)

type service struct {
	repository         Repository
	bantuanRepository  bantuan.Repository
	pembayaranService  pembayaran.Service
}

type Service interface {
	GetTransaksisByBantuanID(input int) ([]Transaksi, error)
	GetTransaksisByUserID(userID int) ([]Transaksi, error)
	CreateTransaksi(input CreateTransaksiInput) (Transaksi, error)
}

func NewService(repository Repository, bantuanRepository bantuan.Repository, pembayaranService pembayaran.Service) *service {
	return &service{repository, bantuanRepository, pembayaranService}
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

func (s *service) CreateTransaksi(input CreateTransaksiInput) (Transaksi, error){
	transaksi := Transaksi{}
	transaksi.BantuanID = input.BantuanID
	transaksi.JumlahUang = input.JumlahUang
	transaksi.UserID	= input.User.ID
	transaksi.StatusTransaksi = "pending"

	newTransaksi, err := s.repository.Save(transaksi)
	if err != nil{
		return newTransaksi, err
	}

	pembayaranTransaksi := pembayaran.Transaksi{
		ID	: newTransaksi.ID,
		JumlahUang : newTransaksi.JumlahUang,
	}

	pembayaranUrl, err := s.pembayaranService.GetPembayaranUrl(pembayaranTransaksi, input.User)
		if err != nil{
			return newTransaksi, err
		}

		newTransaksi.PembayaranUrl = pembayaranUrl

		newTransaksi, err = s.repository.Update(newTransaksi)
		if err != nil{
			return newTransaksi, err
		}

		return newTransaksi, nil
}
