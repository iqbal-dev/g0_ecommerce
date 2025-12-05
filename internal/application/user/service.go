package user

import domain "ecommerce/internal/domain/user"

type Service interface {
	Create(user domain.User) (*domain.User, error)
	Update(id int, user domain.User) (*domain.User, error)
	Delete(id int) (bool, error)
	FindAll() ([]*domain.User, error)
	FindOne(id int) (*domain.User, error)
	FindByEmailAndPassword(email string, password string) (*domain.User, error)
}

type service struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(user domain.User) (*domain.User, error) {
	return s.repo.Create(user)
}

func (s *service) Update(id int, user domain.User) (*domain.User, error) {
	return s.repo.Update(id, user)
}

func (s *service) Delete(id int) (bool, error) {
	return s.repo.Delete(id)
}

func (s *service) FindAll() ([]*domain.User, error) {
	return s.repo.FindAll()
}

func (s *service) FindOne(id int) (*domain.User, error) {
	return s.repo.FindOne(id)
}

func (s *service) FindByEmailAndPassword(email string, password string) (*domain.User, error) {
	return s.repo.FindByEmailAndPassword(email, password)
}
