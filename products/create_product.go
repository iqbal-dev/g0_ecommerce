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

	// Assign a new ID and append the product to the list
	newProduct.Id = len(database.Products) + 1
	database.Products = append(database.Products, newProduct)

	utils.SendJSONResponse(res, http.StatusCreated, "Product created successfully", newProduct)
}
