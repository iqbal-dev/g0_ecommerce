package products

import "ecommerce/domain"

type Service interface {
	Create(product domain.Product) (*domain.Product, error)
	Update(id int, product domain.Product) (*domain.Product, error)
	Delete(id int) (bool, error)
	FindAll() ([]*domain.Product, error)
	FindOne(id int) (*domain.Product, error)
}