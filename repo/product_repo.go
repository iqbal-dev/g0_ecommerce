package repo

import (
	"database/sql"

	"ecommerce/domain"
	"ecommerce/product"

	"github.com/jmoiron/sqlx"
)


type ProductRepo interface {
	product.ProductRepo
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

func (r *productRepo) Create(product domain.Product) (*domain.Product, error) {
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

	var p domain.Product
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

func (r *productRepo) Update(id int, product domain.Product) (*domain.Product, error) {
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
	var p domain.Product
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

func (r *productRepo) FindAll(page, limit int64) ([]*domain.Product, int64, error) {
    // Calculate offset
    offset := (page - 1) * limit

    // Query to fetch paginated products
    query := `
        SELECT 
            id, name, price, description, img_url, created_at, updated_at
        FROM products
        ORDER BY id DESC
        LIMIT $1 OFFSET $2;
    `

    var products []*domain.Product

    // Fetch products
    err := r.db.Select(&products, query, limit, offset)
    if err != nil {
        return nil, 0, err
    }

    // Query to count total products
    countQuery := `SELECT COUNT(*) FROM products;`

    var total int64
    err = r.db.Get(&total, countQuery)
    if err != nil {
        return nil, 0, err
    }

    return products, total, nil
}

func (r *productRepo) FindOne(id int) (*domain.Product, error) {
    query := `
    SELECT 
        id, name, price, description, img_url, created_at, updated_at
    FROM products
    WHERE id = $1;
    `

    var product domain.Product

    err := r.db.Get(&product, query, id)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil // not found
        }
        return nil, err
    }

    return &product, nil
}

