package transaksi

import (
	"time"
	"minipro/user"
	"minipro/bantuan"
)

type Transaksi struct {
	ID         			int
	BantuanID 			int
	UserID     			int
	JumlahUang     		int
	StatusTransaksi     string
	KodeTransaksi       string
	PembayaranUrl		string
	User				user.User
	Bantuan				bantuan.Bantuan
	Created 		 	time.Time
	Updated  			time.Time
}

