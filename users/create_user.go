package users

import (
	"ecommerce/repo"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateUser(res http.ResponseWriter, req *http.Request) {

	var newUser UserType
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&newUser)

	if err != nil {
		utils.SendJSONResponse(res, http.StatusBadRequest, "Invalid request payload", nil)
	}

	user, err := h.userRepo.Create(repo.User{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: newUser.Password,
	})
	if err != nil {
		utils.SendJSONResponse(res, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.SendJSONResponse(res, http.StatusCreated, "User created successfully", user)

}
