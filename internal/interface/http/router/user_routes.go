package router

import (
	app "ecommerce/internal/application/user"
	infradb "ecommerce/internal/infrastructure/db/postgres"
	"ecommerce/internal/interface/http/handlers/users"

	"github.com/jmoiron/sqlx"
)

func RegisterUserRoutes(router *Manager, dbCon *sqlx.DB) {
	userRepo := infradb.NewUserRepository(dbCon)
	userSvc := app.NewService(userRepo)
	handler := users.NewHandler(userSvc)

	router.POST("/login", handler.Login)
	router.POST("/users", handler.CreateUser)
	router.GET("/users", handler.GetUsers)
	router.GET("/users/{id}", handler.GetUser)
	router.PATCH("/users/{id}", handler.UpdateUser)
	router.DELETE("/users/{id}", handler.DeleteUser)
}
