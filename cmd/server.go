package cmd

import (
	"ecommerce/config"
	"ecommerce/middleware"
	"ecommerce/products"
	"ecommerce/routes"
	"fmt"
	"net/http"
)

func Serve(cnf config.Config) {
	fmt.Printf("Type of arg: %T\n", middleware.AuthMiddleware)
	fmt.Printf("Type of arg: %T\n", products.GetProducts)
	router := routes.NewManager(http.NewServeMux())
	routes.RegisterRoutes(router)
	fmt.Println("🚀 Server running on port :", cnf.HttpPort)
	err := http.ListenAndServe(cnf.HttpPort, middleware.Cors(router))
	if err != nil {
		fmt.Println("❌ Error starting server:", err)
	}
}
