package transaksi

type GetBantuanTransaksiInput struct {
	ID int `param:"id" validate:"required"`
}
