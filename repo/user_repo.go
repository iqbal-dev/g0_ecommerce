package repo

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type User struct {
	Id        int    `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"password"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}


type UserRepo interface {
	Create(user User) (*User, error)
	Update(id int, user User) (*User, error)
	Delete(id int) (bool, error)
	FindAll() ([]*User, error)
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

func (r *userRepo) FindAll() ([]*User, error) {
	var users []*User

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


func (r *userRepo) FindOne(id int) (*User, error) {
	var user User

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

func (r *userRepo) FindByEmailAndPassword(email string, password string) (*User, error) {
	var user User

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
