package products

import (
	"ecommerce/utils"
	"fmt"
	"net/http"
	"strconv"
)

// getProducts handles GET requests to the "/products" endpoint.
// It returns a JSON response containing all products.
func (h *Handler) GetProducts(res http.ResponseWriter, req *http.Request) {
		// Default pagination values
	page := 1
	limit := 10

	// Read from query parameters
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

	fmt.Println(page,limit)
	products,total, err := h.svc.FindAll(int64(page), int64(limit))
	if err != nil {
		utils.SendJSONResponse(res, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.SendResponseWithPagination(res,products,int64(page), int64(limit), total)
}
