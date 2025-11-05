package products

import (
	"ecommerce/utils"
	"net/http"
)

// getProducts handles GET requests to the "/products" endpoint.
// It returns a JSON response containing all products.
func (h *Handler) GetProducts(res http.ResponseWriter, req *http.Request) {
	products := h.productRepo.FindALl()
	utils.SendJSONResponse(res, http.StatusOK, "Products retrieved successfully", products)
}
