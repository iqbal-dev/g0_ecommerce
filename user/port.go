package user

import (
	"ecommerce/domain"
	"ecommerce/rest/handlers/users"
)
  type Service interface{
		users.Service
  }


  type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Update(id int, user domain.User) (*domain.User, error)
	Delete(id int) (bool, error)
	FindAll() ([]*domain.User, error)
	FindOne(id int) (*domain.User, error)
	FindByEmailAndPassword(email string, password string) (*domain.User, error)
}