package users

import (
	"ecommerce/domain"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateUser(res http.ResponseWriter, req *http.Request) {

	var newUser domain.User
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&newUser)

	if err != nil {
		utils.SendJSONResponse(res, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	user, err := h.svc.Create(newUser)
	if err != nil {
		utils.SendJSONResponse(res, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.SendJSONResponse(res, http.StatusCreated, "User created successfully", user)

}
