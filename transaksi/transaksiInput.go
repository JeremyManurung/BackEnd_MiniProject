package transaksi

import(
	"minipro/user"
)

type GetBantuanTransaksiInput struct {
	ID 		int `param:"id" validate:"required"`
	User	user.User	
}

type CreateTransaksiInput struct{
	JumlahUang 		int			`json:"jumlah_uang"`
	BantuanID		int			`json:"bantuan_id"`
	User 			user.User
}