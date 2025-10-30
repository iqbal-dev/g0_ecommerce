package products

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func GetProductByID(res http.ResponseWriter, req *http.Request) {

	productId := req.PathValue("id")

	id, err := strconv.Atoi(productId)
	if err != nil {
		utils.SendJSONResponse(res, http.StatusBadRequest, "Invalid product ID", nil)
		return
	}

	for _, product := range database.Products {
		if product.Id == id {
			utils.SendJSONResponse(res, http.StatusOK, "Product retrieved successfully", product)
			return
		}

	}
	utils.SendJSONResponse(res, http.StatusNotFound, "Product not found", nil)
	// Implementation for getting a product by ID will go here
}
