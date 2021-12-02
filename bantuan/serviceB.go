package bantuan

import (
	"fmt"
	"github.com/gosimple/slug"
)

type Service interface {
	FindBantuans(userID int) ([]Bantuan, error)
	CreateBantuan(input CreateBantuanInput) (Bantuan, error)
	FindBantuanByID(input int) (Bantuan, error)
	UpdateBantuan(inputID int, inputData CreateBantuanInput)(Bantuan,error)
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

func (s *service) FindBantuanByID(input int) (Bantuan, error) {
	bantuan, err := s.repository.FindByID(input)

	if err != nil {
		return bantuan, err
	}

	return bantuan, nil
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


func (s *service) UpdateBantuan(inputID int, inputData CreateBantuanInput) (Bantuan, error) {
	bantuan, err := s.repository.FindByID(inputID)
	if err != nil{
		return bantuan,err
	}

	bantuan.TittleBantuan = inputData.TittleBantuan
	bantuan.DeskripsiSingkat = inputData.DeskripsiSingkat
	bantuan.Deskripsi = inputData.Deskripsi
	bantuan.ListKondisi = inputData.ListKondisi
	bantuan.JumlahTarget = inputData.JumlahTarget

	updatedBantuan, err := s.repository.Update(bantuan)
	if err != nil{
		return updatedBantuan,err
	}

	return updatedBantuan, nil

}