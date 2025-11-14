package product

import (
	"ecommerce/domain"
	productHandler "ecommerce/rest/handlers/products"
)

type Service interface{
	productHandler.Service
}
type ProductRepo interface {
	Create(product domain.Product) (*domain.Product, error)
	Update(id int, product domain.Product) (*domain.Product, error)
	Delete(id int) (bool, error)
	FindAll() ([]*domain.Product, error)
	FindOne(id int) (*domain.Product, error)
}
