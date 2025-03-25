package product

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) ListAll() ([]Product, error) {
	return s.repo.GetAll()
}

func (s *Service) Create(product Product) error {
	return s.repo.Create(product)
}

func (s *Service) GetOne(id int) (*Product, error) {
	return s.repo.GetOne(id)
}
