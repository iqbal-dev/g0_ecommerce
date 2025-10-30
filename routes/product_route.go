package routes

import (
	"ecommerce/middleware"
	"ecommerce/products"
)

// RegisterProductRoutes registers all product-related routes
func RegisterProductRoutes(router *Manager) {
	router.GET("/products",
		middleware.ExecutionTimeMiddleware,
		middleware.LoggingMiddleware,
		middleware.AuthMiddleware,
		products.GetProducts,
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
		products.CreateProduct,
	)
	router.PATCH("/products/{id}",
		middleware.LoggingMiddleware,
		middleware.AuthMiddleware,
		products.UpdateProductById,
	)
	router.DELETE("/products/{id}",
		middleware.LoggingMiddleware,
		middleware.AuthMiddleware,
		products.DeleteProductById,
	)
}
