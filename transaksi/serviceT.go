package transaksi

type service struct {
	repository         Repository
}

type Service interface {
	GetTransaksisByBantuanID(input GetBantuanTransaksiInput) ([]Transaksi, error)
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTransaksisByBantuanID(input GetBantuanTransaksiInput) ([]Transaksi, error) {
	transaksis, err := s.repository.GetByBantuanID(input.ID)
	if err != nil {
		return []Transaksi{}, err
	}

	return transaksis, nil
}