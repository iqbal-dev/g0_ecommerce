package users

import "ecommerce/domain"


type Service interface {
	Create(user domain.User) (*domain.User, error)
	Update(id int, user domain.User) (*domain.User, error)
	Delete(id int) (bool, error)
	FindAll() ([]*domain.User, error)
	FindOne(id int) (*domain.User, error)
	FindByEmailAndPassword(email string, password string) (*domain.User, error)
}