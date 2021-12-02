package transaksi

import(
	"minipro/user"
)

type GetBantuanTransaksiInput struct {
	ID 		int `param:"id" validate:"required"`
	User	user.User	
}
