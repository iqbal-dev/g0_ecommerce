package products

import (
	"ecommerce/repo"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type UpdateProduct struct {
	Id          int     `json:"id"`          // Unique identifier for the product
	Name        string  `json:"name"`        // Product name
	Price       float64 `json:"price"`       // Product price
	Description string  `json:"description"` // Product description
	ImgUrl      string  `json:"img_url"`     // Product image URL
}

func (h *Handler) UpdateProductById(res http.ResponseWriter, req *http.Request) {

	productId := req.PathValue("id")

	id, err := strconv.Atoi(productId)
	if err != nil {
		utils.SendJSONResponse(res, http.StatusBadRequest, "Invalid product ID", nil)
		return
	}

	var newProduct UpdateProduct
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&newProduct)

	if err != nil {
		utils.SendJSONResponse(res, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	newProduct.Id = id

	product, err := h.productRepo.Update(id, repo.Product{
		Name:        newProduct.Name,
		Price:       newProduct.Price,
		Description: newProduct.Description,
		ImgUrl:      newProduct.ImgUrl,
	})

	if err != nil {

		utils.SendJSONResponse(res, http.StatusNotFound, "Product not found", product)
		return
	}
	utils.SendJSONResponse(res, http.StatusOK, "Product updated successfully", product)

} // Implementation for getting a product by ID will go here
