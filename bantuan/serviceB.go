package bantuan

type Service interface {
	FindBantuans(userID int) ([]Bantuan, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindBantuans(userID int) ([]Bantuan, error) {
	if userID != 0 {
		bantuans, err := s.repository.FindByUserID(userID)
		if err != nil {
			return bantuans, err
		}

		return bantuans, nil
	}

	bantuans, err := s.repository.FindAll()
	if err != nil {
		return bantuans, err
	}

	return bantuans, nil
}


