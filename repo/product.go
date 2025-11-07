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
	Delete(id int) (bool, error)
	FindALl() ([]*Product, error)
	FindOne(id int) (*Product, error)
}

type productRepo struct {
	productList []*Product
	lastID      int
}

func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateInitialProducts(repo)
	return repo
}

func (r *productRepo) Create(product Product) (*Product, error) {
	r.lastID++
	product.Id = r.lastID
	r.productList = append(r.productList, &product)
	return &product, nil
}

func (r *productRepo) Update(id int, product Product) (*Product, error) {
	for i := range r.productList {
		if r.productList[i].Id == id {
			product.Id = id // ensure original ID
			r.productList[i] = &product
			return &product, nil
		}
	}
	return nil, errors.New("product not found")
}

func (r *productRepo) Delete(id int) (bool, error) {
	var newList []*Product
	deleted := false

	for _, p := range r.productList {
		if p.Id == id {
			deleted = true
			continue
		}
		newList = append(newList, p)
	}

	r.productList = newList
	return deleted, nil
}

func (r *productRepo) FindALl() ([]*Product, error) {
	return r.productList, nil
}

func (r *productRepo) FindOne(id int) (*Product, error) {
	for _, p := range r.productList {
		if p.Id == id {
			return p, nil
		}
	}
	return nil, errors.New("product not found")
}

func generateInitialProducts(r *productRepo) {
	products := []Product{
		{1, "Product 1", 10.99, "This is product 1", "http://example.com/product1.jpg"},
		{2, "Product 2", 15.49, "This is product 2", "http://example.com/product2.jpg"},
		{3, "Product 3", 7.99, "This is product 3", "http://example.com/product3.jpg"},
		{4, "Product 4", 12.75, "This is product 4", "http://example.com/product4.jpg"},
		{5, "Product 5", 9.50, "This is product 5", "http://example.com/product5.jpg"},
		{6, "Product 6", 20.00, "This is product 6", "http://example.com/product6.jpg"},
	}

	for i := range products {
		r.productList = append(r.productList, &products[i])
	}

	r.lastID = len(products)
}
