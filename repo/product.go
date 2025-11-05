package repo

import "errors"

type Product struct {
	Id          int     `json:"id"`          // Unique identifier for the product
	Name        string  `json:"name"`        // Product name
	Price       float64 `json:"price"`       // Product price
	Description string  `json:"description"` // Product description
	ImgUrl      string  `json:"img_url"`     // Product image URL
}

type ProductRepo interface {
	Create(product Product) (*Product, error)
	Update(id int, product Product) (*Product, error)
	Delete(id int) bool
	FindALl() []*Product
	FindOne(id int) *Product
}

type productRepo struct {
	productList []*Product
}

/***
 * * constructor or constructor function
 * * constructor means creator of a object
 */

func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateInitialProducts(repo)
	return repo
}

func (r *productRepo) Create(product Product) (*Product, error) {
	product.Id = len(r.productList) + 1
	r.productList = append(r.productList, &product)
	return &product, nil
}
func (r *productRepo) Update(id int, product Product) (*Product, error) {
	for i := 0; i < len(r.productList); i++ {
		if r.productList[i].Id == id {
			r.productList[i] = &product
			return r.productList[i], nil
		}
	}
	return nil, errors.New("product not found")
}
func (r *productRepo) Delete(id int) bool {
	var tempProductList []*Product
	var isDeleted bool = false

	for i := 0; i < len(r.productList); i++ {
		if r.productList[i].Id != id {
			tempProductList = append(tempProductList, r.productList[i])
		} else {
			isDeleted = true
		}
	}
	r.productList = tempProductList
	return isDeleted
}
func (r *productRepo) FindALl() []*Product {
	return r.productList
}
func (r *productRepo) FindOne(id int) *Product {
	for i := 0; i < len(r.productList); i++ {
		if r.productList[i].Id == id {
			return r.productList[i]
		}
	}
	return nil
}

func generateInitialProducts(r *productRepo) {
	prd1 := Product{1, "Product 1", 10.99, "This is product 1", "http://example.com/product1.jpg"}
	prd2 := Product{2, "Product 2", 15.49, "This is product 2", "http://example.com/product2.jpg"}
	prd3 := Product{3, "Product 3", 7.99, "This is product 3", "http://example.com/product3.jpg"}
	prd4 := Product{4, "Product 4", 12.75, "This is product 4", "http://example.com/product4.jpg"}
	prd5 := Product{5, "Product 5", 9.50, "This is product 5", "http://example.com/product5.jpg"}
	prd6 := Product{6, "Product 6", 20.00, "This is product 6", "http://example.com/product6.jpg"}

	r.productList = append(r.productList, &prd1, &prd2, &prd3, &prd4, &prd5, &prd6)
}
