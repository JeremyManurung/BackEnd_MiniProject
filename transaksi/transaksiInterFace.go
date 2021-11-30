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