package routes

import (
	"ecommerce/products"
	"ecommerce/repo"
	"ecommerce/rest/middleware"
)

// RegisterProductRoutes registers all product-related routes
func RegisterProductRoutes(router *Manager) {
	router.Use(middleware.LoggingMiddleware)
	router.GET("/products",
		middleware.ExecutionTimeMiddleware,
		middleware.LoggingMiddleware,
		middleware.AuthMiddleware,
		products.NewHandler(repo.NewProductRepo()).GetProducts,
	)

	router.GET("/products/{id}", // handle /products/{id} inside handler
		middleware.ExecutionTimeMiddleware,
		middleware.LoggingMiddleware,
		middleware.AuthMiddleware,
		products.GetProductByID,
	)
	router.POST("/products",
		middleware.LoggingMiddleware,
		middleware.AuthMiddleware,
		products.NewHandler(repo.NewProductRepo()).CreateProduct,
	)
	router.PATCH("/products/{id}",
		middleware.LoggingMiddleware,
		middleware.AuthMiddleware,
		products.NewHandler(repo.NewProductRepo()).UpdateProductById,
	)
	router.DELETE("/products/{id}",
		middleware.LoggingMiddleware,
		middleware.AuthMiddleware,
		products.DeleteProductById,
	)
}
