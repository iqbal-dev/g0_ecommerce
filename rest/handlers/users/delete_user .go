package users

import (
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteUser(res http.ResponseWriter, req *http.Request) {

	userId := req.PathValue("id")

	id, err := strconv.Atoi(userId)

	if err != nil {
		utils.SendJSONResponse(res, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	isDeleted, err := h.svc.Delete(id)
	if !isDeleted {
		utils.SendJSONResponse(res, http.StatusNotFound,"User not found", nil)
		return
	}
	utils.SendJSONResponse(res, http.StatusOK, "Users deleted successfully",nil)

}
