package router

import "github.com/jmoiron/sqlx"

func RegisterRoutes(router *Manager, dbCon *sqlx.DB) {
	RegisterProductRoutes(router, dbCon)
	RegisterUserRoutes(router, dbCon)
}
