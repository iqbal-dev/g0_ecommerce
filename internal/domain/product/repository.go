package product

type Repository interface {
	Create(product Product) (*Product, error)
	Update(id int, product Product) (*Product, error)
	Delete(id int) (bool, error)
	FindAll(page, limit int64) ([]*Product, int64, error)
	FindOne(id int) (*Product, error)
}
