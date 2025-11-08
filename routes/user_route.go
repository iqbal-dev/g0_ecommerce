package routes

import (
	"ecommerce/repo"
	"ecommerce/users"

	"github.com/jmoiron/sqlx"
)

func RegisterUserRoutes(router *Manager, dbCon *sqlx.DB) {
	userRepo := repo.NewUserRepo(dbCon)
	handler := users.NewHandler(userRepo)
	router.POST("/login", handler.Login)
	router.POST("/users", handler.CreateUser)
	router.GET("/users", handler.GetUsers)
	router.GET("/users", handler.GetUsers)
	router.GET("/users/{id}", handler.GetUser)
	router.PATCH("/users/{id}", handler.UpdateUser)
	router.DELETE("/users/{id}", handler.DeleteUser)
}
