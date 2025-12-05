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
	var payload LoginType
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&payload); err != nil {
		utils.SendJSONResponse(res, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	user, err := h.svc.FindByEmailAndPassword(payload.Email, payload.Password)
	if err != nil {
		utils.SendJSONResponse(res, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	token, err := utils.GenerateToken(utils.Payload{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	})
	if err != nil {
		utils.SendJSONResponse(res, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.SendJSONResponse(res, http.StatusOK, "User logged in successfully", token)
}
