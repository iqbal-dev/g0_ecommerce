package domain


type Product struct {
	Id          int     `json:"id" db:"id"`          // Unique identifier for the product
	Name        string  `json:"name" db:"name"`        // Product name
	Price       float64 `json:"price" db:"price"`       // Product price
	Description string  `json:"description" db:"description"` // Product description
	ImgUrl      string  `json:"img_url" db:"img_url"`     // Product image URL
	CreatedAt   string  `json:"created_ast" db:"created_at"`
	UpdatedAt   string  `json:"updated_at" db:"updated_at"`
}