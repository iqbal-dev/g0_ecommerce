package user

type Repository interface {
	Create(user User) (*User, error)
	Update(id int, user User) (*User, error)
	Delete(id int) (bool, error)
	FindAll() ([]*User, error)
	FindOne(id int) (*User, error)
	FindByEmailAndPassword(email string, password string) (*User, error)
}
