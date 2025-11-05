package products

import "ecommerce/repo"

type Handler struct {
	// Add any dependencies or services here if needed
	productRepo repo.ProductRepo
}

func NewHandler(productRepo repo.ProductRepo) *Handler {
	return &Handler{
		productRepo: productRepo,
	}
}
