package product
func(svc *service)	Delete(id int) (bool, error){
	return svc.productRepo.Delete(id)
}