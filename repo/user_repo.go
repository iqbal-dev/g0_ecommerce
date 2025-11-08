package repo

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	Update(id int, user User) (*User, error)
	Delete(id int) (bool, error)
	FindALl() ([]*User, error)
	FindOne(id int) (*User, error)
	FindByEmailAndPassword(email string, password string) (*User, error)
}

type userRepo struct {
	db *sqlx.DB
	lastID   int
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	repo := &userRepo{
		db: db,
	}
	return repo
}

func (r *userRepo) Create(user User) (*User, error) {
	query := `
		INSERT INTO users (name, email, password)
		VALUES (:name, :email, :password)
		RETURNING id, name, email, password, created_at, updated_at;
	`

	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var newUser User
	if rows.Next() {
		if err := rows.StructScan(&newUser); err != nil {
			return nil, err
		}
	}

	return &newUser, nil
}

func (r *userRepo) Update(id int, user User) (*User, error) {
		user.Id = id
		query := `
		UPDATE users
			SET name = :name,
				email = :email,
				password = :password
		WHERE id = :id
		RETURNING id, name, email, password, created_at, updated_at;
		`
		rows, err := r.db.NamedQuery(query, user)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		var updatedUser User
		if rows.Next() {
			if err := rows.StructScan(&updatedUser); err != nil {
				return nil, err
			}
		}
		return &updatedUser, nil
	

}

func (r *userRepo) Delete(id int) (bool, error) {
	
}

func (r *userRepo) FindALl() ([]*User, error) {

	if len(r.userList) == 0 {
		return make([]*User, 0), nil
	}
	return r.userList, nil

}

func (r *userRepo) FindOne(id int) (*User, error) {
	for _, u := range r.userList {
		if u.Id == id {
			return u, nil
		}
	}
	return nil, errors.New("user not found")
}
func (r *userRepo) FindByEmailAndPassword(email string, password string) (*User, error) {
	for _, u := range r.userList {
		if u.Email == email && u.Password == password {
			return u, nil
		}
	}
	return nil, errors.New("user not found")
}
