package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/rest/middleware"
	"ecommerce/routes"
	"fmt"
	"net/http"
)

func Serve(cnf *config.Config) {
	fmt.Println("🚀 Starting server...")
	dbCon, err := db.NewConnection()
	if err != nil {
		fmt.Println("❌ Error connecting to database:", err)
		return
	}

	fmt.Println("🚀 Database connected", dbCon.Stats())

	router := routes.NewManager(http.NewServeMux())
	routes.RegisterRoutes(router)
	fmt.Println("🚀 Server running on port :", cnf.HttpPort)
	err = http.ListenAndServe(cnf.HttpPort, middleware.Cors(router))
	if err != nil {
		fmt.Println("❌ Error starting server:", err)
	}
}
