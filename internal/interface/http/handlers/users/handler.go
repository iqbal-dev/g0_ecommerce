package users

import app "ecommerce/internal/application/user"

type Handler struct {
	svc app.Service
}

func NewHandler(svc app.Service) *Handler {
	return &Handler{svc: svc}
}
