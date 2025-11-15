package product

import "ecommerce/domain"



func(svc *service)	Update(id int, product domain.Product) (*domain.Product, error){
	return svc.productRepo.Update(id, product)
}