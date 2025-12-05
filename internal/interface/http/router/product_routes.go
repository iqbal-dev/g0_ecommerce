package router

import (
	app "ecommerce/internal/application/product"
	infradb "ecommerce/internal/infrastructure/db/postgres"
	"ecommerce/internal/interface/http/handlers/products"
	"ecommerce/internal/interface/http/middleware"

	"github.com/jmoiron/sqlx"
)

func RegisterProductRoutes(router *Manager, dbCon *sqlx.DB) {
	productRepo := infradb.NewProductRepository(dbCon)
	productSvc := app.NewService(productRepo)
	handler := products.NewHandler(productSvc)

	router.Use(middleware.LoggingMiddleware)

	router.GET("/products",
		middleware.ExecutionTimeMiddleware,
		middleware.LoggingMiddleware,
		middleware.AuthMiddleware,
		handler.GetProducts,
	)

	router.GET("/products/{id}",
		middleware.ExecutionTimeMiddleware,
		middleware.LoggingMiddleware,
		middleware.AuthMiddleware,
		handler.GetProductByID,
	)

	router.POST("/products",
		middleware.LoggingMiddleware,
		middleware.AuthMiddleware,
		handler.CreateProduct,
	)

	router.PATCH("/products/{id}",
		middleware.LoggingMiddleware,
		middleware.AuthMiddleware,
		handler.UpdateProductById,
	)

	router.DELETE("/products/{id}",
		middleware.LoggingMiddleware,
		middleware.AuthMiddleware,
		handler.DeleteProductById,
	)
}
