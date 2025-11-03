package routes

import (
	"ecommerce/utils"
	"net/http"
)

func RegisterUserRoutes(router *Manager) {

	router.GET("/users", func(res http.ResponseWriter, req *http.Request) {
		utils.SendJSONResponse(res, http.StatusOK, "Users retrieved successfully", nil)
	})
	router.POST("/login", func(res http.ResponseWriter, req *http.Request) {
		token, err := utils.GenerateToken("iqbal")
		if err != nil {
			utils.SendJSONResponse(res, http.StatusInternalServerError, "Failed to generate token", nil)
			return
		}
		utils.SendJSONResponse(res, http.StatusOK, "Users retrieved successfully", token)
	})
}
