package komentar

import (
	"minipro/user"
)

type CreateKomentarInput struct{
	IsiKomentar 	string		`json:"isi_komentar"`
	User 			user.User
}