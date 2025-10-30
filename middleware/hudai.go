package middleware

import (
	"fmt"
	"net/http"
)

func HudaiMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("HudaiMiddleware executed")

		next.ServeHTTP(w, r)
	})

	// This function is intentionally left blank.
}
