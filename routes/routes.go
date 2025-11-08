package routes

import "github.com/jmoiron/sqlx"

// RegisterRoutes registers all modules/routes in the app
func RegisterRoutes(router *Manager, dbCon *sqlx.DB) {
	// Register product routes
	RegisterProductRoutes(router, dbCon)
	RegisterUserRoutes(router, dbCon)

	// You can add more module routes here in the future
	// RegisterUserRoutes(router)
	// RegisterOrderRoutes(router)
}
