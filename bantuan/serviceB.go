package bantuan

import (
	"fmt"
	"github.com/gosimple/slug"
)

type Service interface {
	FindBantuans(userID int) ([]Bantuan, error)
	CreateBantuan(input CreateBantuanInput) (Bantuan, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindBantuans(userID int) ([]Bantuan, error) {
	if userID != 0 {
		bantuans, err := s.repository.FindByUserID(userID)
		if err != nil {
			return bantuans, err
		}

		return bantuans, nil
	}

	bantuans, err := s.repository.FindAll()
	if err != nil {
		return bantuans, err
	}

	return bantuans, nil
}

func (s *service) CreateBantuan(input CreateBantuanInput) (Bantuan, error) {
	bantuan := Bantuan{}
	bantuan.TittleBantuan = input.TittleBantuan
	bantuan.DeskripsiSingkat = input.DeskripsiSingkat
	bantuan.Deskripsi = input.Deskripsi
	bantuan.ListKondisi = input.ListKondisi
	bantuan.JumlahTarget = input.JumlahTarget
	bantuan.UserID = input.User.ID

	slugCandidate := fmt.Sprintf("%s %d", input.TittleBantuan, input.User.ID)
	bantuan.Prm = slug.Make(slugCandidate)

	newBantuan, err := s.repository.Save(bantuan)
	if err != nil {
		return newBantuan, err
	}

	return newBantuan, nil
}




