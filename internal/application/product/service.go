package product

import (
	domain "ecommerce/internal/domain/product"
)

type Service interface {
	Create(product domain.Product) (*domain.Product, error)
	Update(id int, product domain.Product) (*domain.Product, error)
	Delete(id int) (bool, error)
	FindAll(page, limit int64) ([]*domain.Product, int64, error)
	FindOne(id int) (*domain.Product, error)
}

type service struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(product domain.Product) (*domain.Product, error) {
	return s.repo.Create(product)
}

func (s *service) Update(id int, product domain.Product) (*domain.Product, error) {
	return s.repo.Update(id, product)
}

func (s *service) Delete(id int) (bool, error) {
	return s.repo.Delete(id)
}

func (s *service) FindAll(page, limit int64) ([]*domain.Product, int64, error) {
	return s.repo.FindAll(page, limit)
}

func (s *service) FindOne(id int) (*domain.Product, error) {
	return s.repo.FindOne(id)
}
