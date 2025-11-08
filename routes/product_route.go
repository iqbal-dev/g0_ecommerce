package routes

import (
	"ecommerce/products"
	"ecommerce/repo"
	"ecommerce/rest/middleware"

	"github.com/jmoiron/sqlx"
)

// RegisterProductRoutes registers all product-related routes
func RegisterProductRoutes(router *Manager, dbCon *sqlx.DB) {
	// ✅ create repo once
	productRepo := repo.NewProductRepo(dbCon)

	// ✅ create handler once
	handler := products.NewHandler(productRepo)

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
		handler.GetProductByID, // ✅ use shared handler
	)

	router.POST("/products",
		middleware.LoggingMiddleware,
		middleware.AuthMiddleware,
		handler.CreateProduct, // ✅ shared
	)

	router.PATCH("/products/{id}",
		middleware.LoggingMiddleware,
		middleware.AuthMiddleware,
		handler.UpdateProductById, // ✅ shared
	)

	router.DELETE("/products/{id}",
		middleware.LoggingMiddleware,
		middleware.AuthMiddleware,
		handler.DeleteProductById, // ✅ shared
	)
}
