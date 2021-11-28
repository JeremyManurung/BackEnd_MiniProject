package bantuan

type BantuanFormatter struct{
	ID 					int		`json:"id"`
	UserID				int		`json:"user_id"`
	TittleBantuan		string	`json:"tittle_bantuan"`
	DeskripsiSingkat	string	`json:"deskripsi_singkat"`
	ImgUrl				string	`json:"img_url"`
	JumlahTarget		int		`json:"jumlah_target"`
	JumlahBar			int		`json:"jumlah_bar"`
	Prm					string	`json:"prm"`
}

func FormatBantuan(bantuan Bantuan) BantuanFormatter {
	bantuanFormatter := BantuanFormatter{}
	bantuanFormatter.ID = bantuan.ID
	bantuanFormatter.UserID = bantuan.UserID
	bantuanFormatter.TittleBantuan = bantuan.TittleBantuan
	bantuanFormatter.DeskripsiSingkat = bantuan.DeskripsiSingkat
	bantuanFormatter.JumlahTarget = bantuan.JumlahTarget
	bantuanFormatter.JumlahBar = bantuan.JumlahBar
	bantuanFormatter.Prm = bantuan.Prm
	bantuanFormatter.ImgUrl = ""

	if len(bantuan.BantuanImgs) > 0 {
		bantuanFormatter.ImgUrl  = bantuan.BantuanImgs[0].TittleImg
	}


	return bantuanFormatter
}

func FormatBantuans(bantuans []Bantuan) []BantuanFormatter {
	bantuansFormatter := []BantuanFormatter{}

	for _, bantuan := range bantuans {
		bantuanFormatter := FormatBantuan(bantuan)
		bantuansFormatter = append(bantuansFormatter, bantuanFormatter)
	}

	return bantuansFormatter
}