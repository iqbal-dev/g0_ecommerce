package product

import "ecommerce/domain"


type service struct {
 productRepo	ProductRepo
}	


func NewService(prdctRepo ProductRepo) Service{
	return &service{
		productRepo: prdctRepo,
	}
}



func(svc *service)	Create(product domain.Product) (*domain.Product, error){
	return svc.productRepo.Create(product)
}
func(svc *service)	Update(id int, product domain.Product) (*domain.Product, error){
	return svc.productRepo.Update(id, product)
}
func(svc *service)	Delete(id int) (bool, error){
	return svc.productRepo.Delete(id)
}
func(svc *service)	FindAll() ([]*domain.Product, error){
	return svc.productRepo.FindAll()
}
func(svc *service)	FindOne(id int) (*domain.Product, error){
	return svc.productRepo.FindOne(id)
}