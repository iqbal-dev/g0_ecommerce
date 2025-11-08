package products

import (
	"ecommerce/repo"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

type Product struct { // Unique identifier for the product
	Name        string  `json:"name"`        // Product name
	Price       float64 `json:"price"`       // Product price
	Description string  `json:"description"` // Product description
	ImgUrl      string  `json:"img_url"`     // Product image URL
}

// createProduct handles POST requests to the "/create-product" endpoint.
// It creates a new product based on the JSON body and adds it to the products slice.
func (h *Handler) CreateProduct(res http.ResponseWriter, req *http.Request) {

	// Decode the request body into a Product struct
	var newProduct Product
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		utils.SendJSONResponse(res, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	product, err := h.productRepo.Create(repo.Product{
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
