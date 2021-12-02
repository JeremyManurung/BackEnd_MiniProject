package komentar

import (
	
)

type service struct {
	repository         Repository
}

type Service interface {
	CreateKomentar(input CreateKomentarInput) (Komentar, error)
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateKomentar(input CreateKomentarInput) (Komentar, error) {
	komentar := Komentar{}
	komentar.IsiKomentar = input.IsiKomentar
	komentar.UserID = input.User.ID

	newKomentar, err := s.repository.Save(komentar)
	if err != nil {
		return newKomentar, err
	}

	return newKomentar, nil
}