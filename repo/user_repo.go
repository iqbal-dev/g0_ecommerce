package repo

import (
	"ecommerce/domain"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)




type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Update(id int, user domain.User) (*domain.User, error)
	Delete(id int) (bool, error)
	FindAll() ([]*domain.User, error)
	FindOne(id int) (*domain.User, error)
	FindByEmailAndPassword(email string, password string) (*domain.User, error)
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

func (r *userRepo) Create(user domain.User) (*domain.User, error) {
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

	var newUser domain.User
	if rows.Next() {
		if err := rows.StructScan(&newUser); err != nil {
			return nil, err
		}
	}

	return &newUser, nil
}

func (r *userRepo) Update(id int, user domain.User) (*domain.User, error) {
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

	var updatedUser domain.User
	if rows.Next() {
		// Scan the returned row into updatedUser
		if err := rows.StructScan(&updatedUser); err != nil {
			return nil, err
		}
		return &updatedUser, nil
	}

	// If no row returned, the user ID does not exist
	return nil, errors.New("user not found")
}


func (r *userRepo) Delete(id int) (bool, error) {
	query := `DELETE FROM users WHERE id = :id`

	params := map[string]interface{}{
		"id": id,
	}

	result, err := r.db.NamedExec(query, params)
	fmt.Println(err)
	if err != nil {
		return false, err
	}

	rows, _ := result.RowsAffected()
	return rows > 0, nil
}

func (r *userRepo) FindAll() ([]*domain.User, error) {
	var users []*domain.User

	query := `
		SELECT id, name, email, password, created_at, updated_at
		FROM users
		ORDER BY id DESC
	`

	err := r.db.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}


func (r *userRepo) FindOne(id int) (*domain.User, error) {
	var user domain.User

	query := `
		SELECT id, name, email, password, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	err := r.db.Get(&user, query, id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (r *userRepo) FindByEmailAndPassword(email string, password string) (*domain.User, error) {
	var user domain.User

	query := `
		SELECT id, name, email, password, created_at, updated_at
		FROM users
		WHERE email = $1 AND password = $2
	`

	err := r.db.Get(&user, query, email,password)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}
