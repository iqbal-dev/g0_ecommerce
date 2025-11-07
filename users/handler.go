package users

import "ecommerce/repo"

type UserType struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Handler struct {
	userRepo repo.UserRepo
}

func NewHandler(userRepo repo.UserRepo) *Handler {
	return &Handler{
		userRepo: userRepo,
	}
}
