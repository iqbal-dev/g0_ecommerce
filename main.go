package main

import (
	"ecommerce/cmd"
	"ecommerce/config"
)

// main initializes the HTTP server and registers all the routes.
func main() {
	cnf := config.GetConfig()
	cmd.Serve(cnf)
}

// init populates the products slice with some initial sample data.
