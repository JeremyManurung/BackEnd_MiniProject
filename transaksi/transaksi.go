package transaksi

import (
	"time"
	"minipro/user"
)

type Transaksi struct {
	ID         			int
	BantuanID 			int
	UserID     			int
	JumlahUang     		int
	StatusTransaksi     string
	CodeTransaksi       string
	User				user.User
	Created 		 	time.Time
	Updated  			time.Time
}

