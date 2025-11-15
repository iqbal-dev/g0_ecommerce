package routes

import (
	"ecommerce/repo"
	"ecommerce/rest/handlers/users"
	"ecommerce/user"

	"github.com/jmoiron/sqlx"
)

func RegisterUserRoutes(router *Manager, dbCon *sqlx.DB) {
	userRepo := repo.NewUserRepo(dbCon)
	userSvc := user.NewUserService(userRepo)
	handler := users.NewHandler(userSvc)
	router.POST("/login", handler.Login)
	router.POST("/users", handler.CreateUser)
	router.GET("/users", handler.GetUsers)
	router.GET("/users", handler.GetUsers)
	router.GET("/users/{id}", handler.GetUser)
	router.PATCH("/users/{id}", handler.UpdateUser)
	router.DELETE("/users/{id}", handler.DeleteUser)
}
