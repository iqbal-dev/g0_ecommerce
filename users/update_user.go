package users

import (
	"ecommerce/repo"
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

	var updateProduct UserType
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&updateProduct)
	if err != nil {
		utils.SendJSONResponse(res, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	product, err := h.userRepo.Update(id, repo.User{
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
