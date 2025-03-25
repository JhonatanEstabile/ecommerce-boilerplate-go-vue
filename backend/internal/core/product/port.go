package product

type Repository interface {
	GetAll() ([]Product, error)
	Create(Product) error
	GetOne(int) (*Product, error)
}
