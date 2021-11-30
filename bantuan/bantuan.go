package bantuan

import (
	"time"
	"minipro/user"
)

type Bantuan struct{
	ID					int 
	JumlahBar 			int
	Deskripsi 			string
	DeskripsiSingkat 	string
	ListKondisi 		string
	JumlahPendonasi 	int
	UserID 				int
	TittleBantuan 		string
	Prm					string
	JumlahTarget 		int
	Created				time.Time
	Updated				time.Time
	BantuanImgs			[]BantuanImg
	User				user.User
}

type BantuanImg struct {
	ID 					int
	BantuanID			int
	TittleImg			string
	ImgUtama			int
	Created				time.Time
	Updated				time.Time
}