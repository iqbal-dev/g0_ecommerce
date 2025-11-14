package user

import "ecommerce/domain"


type service struct{
	userRepo UserRepo
}


func NewUserService(userRepo UserRepo)Service{
return &service{
userRepo:userRepo,
}
}

func(svc *service)	Create(user domain.User) (*domain.User, error){
	return svc.userRepo.Create(user)
}
func(svc *service)	Update(id int, user domain.User) (*domain.User, error){
	return svc.userRepo.Update(id,user)
}
func(svc *service)	Delete(id int) (bool, error){
	return svc.userRepo.Delete(id)
}
func(svc *service)	FindAll() ([]*domain.User, error){
	return svc.userRepo.FindAll()
}
func(svc *service)	FindOne(id int) (*domain.User, error){
	return svc.userRepo.FindOne(id)
}
func(svc *service)	FindByEmailAndPassword(email string, password string) (*domain.User, error){
	return svc.userRepo.FindByEmailAndPassword(email, password)
}



