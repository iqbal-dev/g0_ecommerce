package products

type Handler struct {
	// Add any dependencies or services here if needed
	svc Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{
		svc: svc,
	}
}
