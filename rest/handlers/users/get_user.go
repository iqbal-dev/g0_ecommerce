package users

import (
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetUser(res http.ResponseWriter, req *http.Request) {

	userId := req.PathValue("id")

	id, err := strconv.Atoi(userId)

	if err != nil {
		utils.SendJSONResponse(res, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	users, err := h.svc.FindOne(id)
	if err != nil {
		utils.SendJSONResponse(res, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.SendJSONResponse(res, http.StatusOK, "Users retrieved successfully", users)

}
