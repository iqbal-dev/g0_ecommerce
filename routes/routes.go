package routes

// RegisterRoutes registers all modules/routes in the app
func RegisterRoutes(router *Manager) {
	// Register product routes
	RegisterProductRoutes(router)
	RegisterUserRoutes(router)

	// You can add more module routes here in the future
	// RegisterUserRoutes(router)
	// RegisterOrderRoutes(router)
}
