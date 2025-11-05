package middleware

import (
	"ecommerce/utils"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			utils.SendJSONResponse(w, http.StatusUnauthorized, "Missing authorization header", nil)
			return
		}

		if len(strings.Split(authHeader, " ")) != 2 {
			utils.SendJSONResponse(w, http.StatusUnauthorized, "Invalid authorization header", nil)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		validate, err := utils.VerifyToken(token)
		if !validate {
			utils.SendJSONResponse(w, http.StatusUnauthorized, err, nil)
			return
		}

		r.Header.Set("username", "iqbal")

		next.ServeHTTP(w, r)
	})

	// This function can be used to initialize global routes or middleware if needed in the future.
}
