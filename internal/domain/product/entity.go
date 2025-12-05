package product

type Product struct {
	Id          int     `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Price       float64 `json:"price" db:"price"`
	Description string  `json:"description" db:"description"`
	ImgUrl      string  `json:"img_url" db:"img_url"`
	CreatedAt   string  `json:"created_at" db:"created_at"`
	UpdatedAt   string  `json:"updated_at" db:"updated_at"`
}
