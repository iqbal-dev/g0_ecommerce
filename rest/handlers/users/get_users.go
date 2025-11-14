package users

import (
	"ecommerce/utils"
	"net/http"
)

func (h *Handler) GetUsers(res http.ResponseWriter, req *http.Request) {
	users, err := h.svc.FindAll()
	if err != nil {
		utils.SendJSONResponse(res, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.SendJSONResponse(res, http.StatusOK, "Users retrieved successfully", users)

}
