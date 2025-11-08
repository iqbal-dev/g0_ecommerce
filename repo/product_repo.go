package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	Id          int     `json:"id" db:"id"`          // Unique identifier for the product
	Name        string  `json:"name" db:"name"`        // Product name
	Price       float64 `json:"price" db:"price"`       // Product price
	Description string  `json:"description" db:"description"` // Product description
	ImgUrl      string  `json:"img_url" db:"img_url"`     // Product image URL
	CreatedAt   string  `json:"created_ast" db:"created_at"`
	UpdatedAt   string  `json:"updated_at" db:"updated_at"`
}

type ProductRepo interface {
	Create(product Product) (*Product, error)
	Update(id int, product Product) (*Product, error)
	Delete(id int) (bool, error)
	FindAll() ([]*Product, error)
	FindOne(id int) (*Product, error)
}

type productRepo struct {
	db *sqlx.DB
	lastID      int
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	repo := &productRepo{
		db: db,
	}
	return repo
}

func (r *productRepo) Create(product Product) (*Product, error) {
	query := `
	INSERT INTO products(
		name,
		price,
		description,
		img_url
	) VALUES(
		$1, $2, $3, $4

	) RETURNING id, name, price, description, img_url, created_at, updated_at
	`

	var p Product
	err := r.db.QueryRow(
		query, 
		product.Name, 
		product.Price, 
		product.Description, 
		product.ImgUrl,
		).Scan(
		&p.Id,
		&p.Name,
		&p.Price,
		&p.Description,
		&p.ImgUrl,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	if err != nil{
		return nil,err
	}

	return &p,nil


}

func (r *productRepo) Update(id int, product Product) (*Product, error) {
	query := `
		UPDATE products SET 
			name = $1,
			price = $2,
			description = $3,
			img_url = $4,
			updated_at = NOW()
		WHERE id = $5
		RETURNING id, name, price, description, img_url, created_at, updated_at
	`
	var p Product
	err := r.db.QueryRow(
		query,product.Name, 
		product.Price, 
		product.Description, 
		product.ImgUrl,
		id,
		).Scan(
		&p.Id,
		&p.Name,
		&p.Price,
		&p.Description,
		&p.ImgUrl,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	if err !=nil{
		return nil,err
	}
	return &p,nil



}

func (r *productRepo) Delete(id int) (bool, error) {
    query := `
    DELETE FROM products
    WHERE id = $1
    RETURNING id;
    `

    var deletedId int

    err := r.db.QueryRow(query, id).Scan(&deletedId)
    if err != nil {
        // If no rows matched, it's NOT an error — it's just not found
        if err == sql.ErrNoRows {
            return false, nil
        }
        return false, err
    }

    return true, nil
}

func (r *productRepo) FindAll() ([]*Product, error) {
    query := `
    SELECT 
        id, name, price, description, img_url, created_at, updated_at
    FROM products
    ORDER BY id DESC;
    `

    var products []*Product

    err := r.db.Select(&products, query)
    if err != nil {
        return nil, err
    }

    return products, nil
}

func (r *productRepo) FindOne(id int) (*Product, error) {
    query := `
    SELECT 
        id, name, price, description, img_url, created_at, updated_at
    FROM products
    WHERE id = $1;
    `

    var product Product

    err := r.db.Get(&product, query, id)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil // not found
        }
        return nil, err
    }

    return &product, nil
}

