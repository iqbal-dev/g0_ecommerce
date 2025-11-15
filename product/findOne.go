package product

import "ecommerce/domain"


func(svc *service)	FindOne(id int) (*domain.Product, error){
	return svc.productRepo.FindOne(id)
}