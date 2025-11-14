package products

import (
	"ecommerce/domain"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

// createProduct handles POST requests to the "/create-product" endpoint.
// It creates a new product based on the JSON body and adds it to the products slice.
func (h *Handler) CreateProduct(res http.ResponseWriter, req *http.Request) {

	// Decode the request body into a Product struct
	var newProduct domain.Product
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		utils.SendJSONResponse(res, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	product, err := h.svc.Create(domain.Product{
		Name:        newProduct.Name,
		Price:       newProduct.Price,
		Description: newProduct.Description,
		ImgUrl:      newProduct.ImgUrl,
	})

	if err != nil {
		utils.SendJSONResponse(res, http.StatusInternalServerError, err.Error(), nil)
		return

	}

	utils.SendJSONResponse(res, http.StatusCreated, "Product created successfully", product)
}
