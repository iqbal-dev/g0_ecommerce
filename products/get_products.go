package products

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
)

// getProducts handles GET requests to the "/products" endpoint.
// It returns a JSON response containing all products.
func GetProducts(res http.ResponseWriter, req *http.Request) {

	utils.SendJSONResponse(res, http.StatusOK, "Products retrieved successfully", database.Products)
}
