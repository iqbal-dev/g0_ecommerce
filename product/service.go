package product


type service struct {
 productRepo	ProductRepo
}	


func NewService(prdctRepo ProductRepo) Service{
	return &service{
		productRepo: prdctRepo,
	}
}



