package products

import (
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetProducts(res http.ResponseWriter, req *http.Request) {
	page := 1
	limit := 10

	query := req.URL.Query()

	if p := query.Get("page"); p != "" {
		if v, err := strconv.Atoi(p); err == nil && v > 0 {
			page = v
		}
	}

	if l := query.Get("limit"); l != "" {
		if v, err := strconv.Atoi(l); err == nil && v > 0 {
			limit = v
		}
	}

	products, total, err := h.svc.FindAll(int64(page), int64(limit))
	if err != nil {
		utils.SendJSONResponse(res, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.SendResponseWithPagination(res, products, int64(page), int64(limit), total)
}
