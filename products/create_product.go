package products

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

// createProduct handles POST requests to the "/create-product" endpoint.
// It creates a new product based on the JSON body and adds it to the products slice.
func CreateProduct(res http.ResponseWriter, req *http.Request) {

	// Decode the request body into a Product struct
	var newProduct database.Product
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		utils.SendJSONResponse(res, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	product := database.Create(newProduct)

	utils.SendJSONResponse(res, http.StatusCreated, "Product created successfully", product)
}
