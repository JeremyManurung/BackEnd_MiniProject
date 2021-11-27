package bantuan

import ("time")

type Bantuan struct{
	ID					int
	JumlahBar 			int
	Deskripsi 			string
	DeskripsiSingkat 	string
	ListKondisi 		string
	JumlahPendonasi 	int
	UserID 				int
	TittleBantuan 		string
	Param				string
	JumlahTarget 		int
	Created				time.Time
	Updated				time.Time
	BantuanImgs			[]BantuanImg
}

type BantuanImg struct {
	ID 					int
	BantuanID			int
	TittleImg			string
	ImgUtama			int
	Created				time.Time
	Updated				time.Time
}