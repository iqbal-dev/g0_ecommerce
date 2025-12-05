package users

import (
	domain "ecommerce/internal/domain/user"
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

	var updateUser domain.User
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&updateUser); err != nil {
		utils.SendJSONResponse(res, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	user, err := h.svc.Update(id, domain.User{
		Name:     updateUser.Name,
		Email:    updateUser.Email,
		Password: updateUser.Password,
	})

	if err != nil {
		utils.SendJSONResponse(res, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.SendJSONResponse(res, http.StatusOK, "User updated successfully", user)
}
