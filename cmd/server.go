package cmd

import (
	"ecommerce/config"
	infradb "ecommerce/internal/infrastructure/db/postgres"
	"ecommerce/internal/interface/http/middleware"
	"ecommerce/internal/interface/http/router"
	"fmt"
	"net/http"
)

func Serve(cnf *config.Config) {
	fmt.Println("🚀 Starting server...")
	dbCon, err := infradb.NewConnection(&cnf.DBConfig)
	if err != nil {
		fmt.Println("❌ Error connecting to database:", err)
		return
	}

	fmt.Println("🚀 Database connected")

	r := router.NewManager(http.NewServeMux())
	router.RegisterRoutes(r, dbCon)

	fmt.Println("🚀 Server running on port ", cnf.HttpPort)
	if err := http.ListenAndServe(cnf.HttpPort, middleware.Cors(r)); err != nil {
		fmt.Println("❌ Error starting server:", err)
	}
}
