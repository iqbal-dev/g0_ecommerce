package users

import (
	"ecommerce/domain"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateUser(res http.ResponseWriter, req *http.Request) {
	userId := req.PathValue("id")

	id, err := strconv.Atoi(userId)
	if err != nil {
		utils.SendJSONResponse(res, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	var updateProduct domain.User
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&updateProduct)
	if err != nil {
		utils.SendJSONResponse(res, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	product, err := h.svc.Update(id, domain.User{
		Name:     updateProduct.Name,
		Email:    updateProduct.Email,
		Password: updateProduct.Password,
	})

	if err != nil {
		utils.SendJSONResponse(res, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.SendJSONResponse(res, http.StatusOK, "User updated successfully", product)

}
