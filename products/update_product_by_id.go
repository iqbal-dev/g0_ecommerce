package products

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func UpdateProductById(res http.ResponseWriter, req *http.Request) {

	productId := req.PathValue("id")

	id, err := strconv.Atoi(productId)
	if err != nil {
		utils.SendJSONResponse(res, http.StatusBadRequest, "Invalid product ID", nil)
		return
	}

	// Decode request body into a Product struct
	defer req.Body.Close()
	var updatedProduct database.Product
	if err := json.NewDecoder(req.Body).Decode(&updatedProduct); err != nil {
		utils.SendJSONResponse(res, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	// Ensure the product ID is preserved/consistent
	updatedProduct.Id = id

	for i, product := range database.Products {
		if product.Id == id {
			database.Products[i] = updatedProduct

			utils.SendJSONResponse(res, http.StatusOK, "Product updated successfully", database.Products[i])
			return
		}
	}
	utils.SendJSONResponse(res, http.StatusNotFound, "Product not found", nil)
} // Implementation for getting a product by ID will go here
