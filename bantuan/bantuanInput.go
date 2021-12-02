package bantuan

import(
	"minipro/user"
)

type GetBantuanDetailInput struct {
	ID int `param:"id" validate:"required"`
}

type CreateBantuanInput struct {
	TittleBantuan             string `json:"tittle_bantuan" validate:"required"`
	DeskripsiSingkat 		  string `json:"deskripsi_singkat" validate:"required"`
	Deskripsi   	   		  string `json:"deskripsi" validate:"required"`
	JumlahTarget       	 	  int    `json:"jumlah_target" validate:"required"`
	ListKondisi            	  string `json:"list_kondisi" validate:"required"`
	User             		  user.User
}