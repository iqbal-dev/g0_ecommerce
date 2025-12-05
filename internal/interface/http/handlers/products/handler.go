package products

import app "ecommerce/internal/application/product"

type Handler struct {
	svc app.Service
}

func NewHandler(svc app.Service) *Handler {
	return &Handler{svc: svc}
}
