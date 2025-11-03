package products

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func DeleteProductById(res http.ResponseWriter, req *http.Request) {
	productId := req.PathValue("id")

	id, err := strconv.Atoi(productId)
	if err != nil {
		utils.SendJSONResponse(res, http.StatusBadRequest, "Invalid product ID", nil)
		return
	}

	isDeleted := database.Delete(id)

	if isDeleted {
		utils.SendJSONResponse(res, http.StatusOK, "Product deleted successfully", nil)
		return
	}

	utils.SendJSONResponse(res, http.StatusNotFound, "Product not found", nil)

	//here implement product delete

} // Implementation for getting a product by ID will go here
