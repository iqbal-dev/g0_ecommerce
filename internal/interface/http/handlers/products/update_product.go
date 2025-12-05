package products

import (
	domain "ecommerce/internal/domain/product"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type UpdateProduct struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImgUrl      string  `json:"img_url"`
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
	if err := decoder.Decode(&newProduct); err != nil {
		utils.SendJSONResponse(res, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	product, err := h.svc.Update(id, domain.Product{
		Name:        newProduct.Name,
		Price:       newProduct.Price,
		Description: newProduct.Description,
		ImgUrl:      newProduct.ImgUrl,
	})

	if err != nil {
		utils.SendJSONResponse(res, http.StatusNotFound, err.Error(), product)
		return
	}
	utils.SendJSONResponse(res, http.StatusOK, "Product updated successfully", product)
}
