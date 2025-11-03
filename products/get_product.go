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

	product := database.FindOne(id)
	utils.SendJSONResponse(res, http.StatusNotFound, "Product not found", product)
	// Implementation for getting a product by ID will go here
}
