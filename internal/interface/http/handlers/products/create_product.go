package products

import (
	domain "ecommerce/internal/domain/product"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateProduct(res http.ResponseWriter, req *http.Request) {
	var newProduct domain.Product
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&newProduct); err != nil {
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
