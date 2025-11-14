package users

import (
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

type LoginType struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(res http.ResponseWriter, req *http.Request) {

	var newUser LoginType
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&newUser)

	if err != nil {
		utils.SendJSONResponse(res, http.StatusBadRequest, "Invalid request payload", nil)
	}

	users, err := h.svc.FindByEmailAndPassword(newUser.Email, newUser.Password)
	if err != nil {
		utils.SendJSONResponse(res, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	token, err := utils.GenerateToken(utils.Payload{
		Id:    users.Id,
		Name:  users.Name,
		Email: users.Email,
	})
	if err != nil {
		utils.SendJSONResponse(res, http.StatusInternalServerError, err.Error(), nil)
	}
	utils.SendJSONResponse(res, http.StatusOK, "User logged in successfully", token)

}
