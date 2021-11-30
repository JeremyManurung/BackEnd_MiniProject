package transaksi

import (
	"time"
)

type BantuanTransaksiFormatter struct {
	ID        	 int       `json:"id"`
	Nama      	 string    `json:"nama"`
	JumlahUang   int       `json:"jumlah_uang"`
	Created		 time.Time `json:"created"`
}

func FormatBantuanTransaksi(transaksi Transaksi) BantuanTransaksiFormatter {
	formatter := BantuanTransaksiFormatter{}
	formatter.ID = transaksi.ID
	formatter.Nama = transaksi.User.Nama
	formatter.JumlahUang = transaksi.JumlahUang
	formatter.Created = transaksi.Created
	return formatter
}

func FormatBantuanTransaksis(transaksis []Transaksi) []BantuanTransaksiFormatter {
	if len(transaksis) == 0 {
		return []BantuanTransaksiFormatter{}
	}

	var transaksisFormatter []BantuanTransaksiFormatter

	for _, transaksi := range transaksis {
		formatter := FormatBantuanTransaksi(transaksi)
		transaksisFormatter = append(transaksisFormatter, formatter)
	}

	return transaksisFormatter
}

type UserTransaksiFormatter struct {
	ID        	 	  int       	   `json:"id"`
	JumlahUang        int    		   `json:"jumlah_uang"`
	StatusTransaksi   string     	   `json:"status_transaksi"`
	Created		 	  time.Time  	   `json:"created"`
	Bantuan			  BantuanFormatter `json:"bantuan"`
}

type BantuanFormatter struct {
	TittleBantuan    string      `json:"tittle_bantuan"`
	ImgUrl			 string		 `json:"img_url"`
}

func FormatUserTransaksi(transaksi Transaksi) UserTransaksiFormatter{
	formatter := UserTransaksiFormatter{}
	formatter.ID = transaksi.ID
	formatter.JumlahUang = transaksi.JumlahUang
	formatter.StatusTransaksi = transaksi.StatusTransaksi
	formatter.Created = transaksi.Created


	bantuanFormatter := BantuanFormatter{}
	bantuanFormatter.TittleBantuan = transaksi.Bantuan.TittleBantuan

	bantuanFormatter.ImgUrl = ""
	if len(transaksi.Bantuan.BantuanImgs) > 0{
		bantuanFormatter.ImgUrl = transaksi.Bantuan.BantuanImgs[0].TittleImg
	}
	
	formatter.Bantuan = bantuanFormatter

	return formatter
}

func FormatUserTransaksis(transaksis []Transaksi) []UserTransaksiFormatter {
	if len(transaksis) == 0 {
		return []UserTransaksiFormatter{}
	}

	var transaksisFormatter []UserTransaksiFormatter

	for _, transaksi := range transaksis {
		formatter := FormatUserTransaksi(transaksi)
		transaksisFormatter = append(transaksisFormatter, formatter)
	}

	return transaksisFormatter
}