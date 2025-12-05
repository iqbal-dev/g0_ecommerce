package products

import (
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProductById(res http.ResponseWriter, req *http.Request) {
	productId := req.PathValue("id")

	id, err := strconv.Atoi(productId)
	if err != nil {
		utils.SendJSONResponse(res, http.StatusBadRequest, "Invalid product ID", nil)
		return
	}

	deleted, err := h.svc.Delete(id)
	if err != nil {
		utils.SendJSONResponse(res, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	if !deleted {
		utils.SendJSONResponse(res, http.StatusNotFound, "Product not found", nil)
		return
	}

	utils.SendJSONResponse(res, http.StatusOK, "Product deleted", nil)
}
