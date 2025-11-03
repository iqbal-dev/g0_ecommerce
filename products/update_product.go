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

	var newProduct database.Product
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&newProduct)

	if err != nil {
		utils.SendJSONResponse(res, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	newProduct.Id = id

	product := database.Update(id, newProduct)

	if product == nil {

		utils.SendJSONResponse(res, http.StatusNotFound, "Product not found", product)
	}
	utils.SendJSONResponse(res, http.StatusOK, "Product updated successfully", product)

} // Implementation for getting a product by ID will go here
