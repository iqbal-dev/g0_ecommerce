package database

// Product represents a product entity with essential details.
type Product struct {
	Id          int     `json:"id"`          // Unique identifier for the product
	Name        string  `json:"name"`        // Product name
	Price       float64 `json:"price"`       // Product price
	Description string  `json:"description"` // Product description
	ImgUrl      string  `json:"img_url"`     // Product image URL
}

// products holds all the available products in memory.
var Products []Product

func init() {
	prd1 := Product{1, "Product 1", 10.99, "This is product 1", "http://example.com/product1.jpg"}
	prd2 := Product{2, "Product 2", 15.49, "This is product 2", "http://example.com/product2.jpg"}
	prd3 := Product{3, "Product 3", 7.99, "This is product 3", "http://example.com/product3.jpg"}
	prd4 := Product{4, "Product 4", 12.75, "This is product 4", "http://example.com/product4.jpg"}
	prd5 := Product{5, "Product 5", 9.50, "This is product 5", "http://example.com/product5.jpg"}
	prd6 := Product{6, "Product 6", 20.00, "This is product 6", "http://example.com/product6.jpg"}

	Products = append(Products, prd1, prd2, prd3, prd4, prd5, prd6)
}
