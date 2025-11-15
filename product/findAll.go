package product

import "ecommerce/domain"
func(svc *service)	FindAll(page,limit int64) ([]*domain.Product, int64, error){
	return svc.productRepo.FindAll(page, limit)
}