package products

import (
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetProductByID(res http.ResponseWriter, req *http.Request) {

	productId := req.PathValue("id")

	id, err := strconv.Atoi(productId)
	if err != nil {
		utils.SendJSONResponse(res, http.StatusBadRequest, "Invalid product ID", nil)
		return
	}

	product, err := h.productRepo.FindOne(id)
	if err != nil {
		utils.SendJSONResponse(res, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.SendJSONResponse(res, http.StatusOK, "Product fetch successfully", product)
	// Implementation for getting a product by ID will go here
}
