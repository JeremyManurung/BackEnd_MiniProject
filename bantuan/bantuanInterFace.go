package bantuan

import(
	"strings"
)

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

type BantuanDetailFormatter struct {
	ID               	int                      `json:"id"`
	TittleBantuan       string                   `json:"tittle_bantuan"`
	DeskripsiSingkat 	string                   `json:"deskripsi_singkat"`
	Deskripsi      		string                   `json:"deskripsi"`
	ImgUrl        	 	string                   `json:"img_url"`
	JumlahTarget      	int                      `json:"jumlah_target"`
	JumlahBar    		int                      `json:"jumlah_bar"`
	JumlahPendonasi     int                      `json:"jumlah_pendonasi"`
	UserID           	int                      `json:"user_id"`
	Prm             	string                   `json:"prm"`
	ListKondisi         []string                 `json:"list_kondisi"`
}


func FormatBantuanDetail(bantuan Bantuan) BantuanDetailFormatter {
	bantuanDetailFormatter := BantuanDetailFormatter{}
	bantuanDetailFormatter.ID = bantuan.ID
	bantuanDetailFormatter.TittleBantuan = bantuan.TittleBantuan
	bantuanDetailFormatter.DeskripsiSingkat = bantuan.DeskripsiSingkat
	bantuanDetailFormatter.Deskripsi = bantuan.Deskripsi
	bantuanDetailFormatter.JumlahTarget = bantuan.JumlahTarget
	bantuanDetailFormatter.JumlahBar = bantuan.JumlahBar
	bantuanDetailFormatter.JumlahPendonasi = bantuan.JumlahPendonasi
	bantuanDetailFormatter.UserID = bantuan.UserID
	bantuanDetailFormatter.Prm = bantuan.Prm
	bantuanDetailFormatter.ImgUrl = ""

	if len(bantuan.BantuanImgs) > 0 {
		bantuanDetailFormatter.ImgUrl = bantuan.BantuanImgs[0].TittleImg
	}

	var listkondisis []string

	for _, listkondisi := range strings.Split(bantuan.ListKondisi, ",") {
		listkondisis = append(listkondisis, strings.TrimSpace(listkondisi))
	}

	bantuanDetailFormatter.ListKondisi = listkondisis

	return bantuanDetailFormatter
}