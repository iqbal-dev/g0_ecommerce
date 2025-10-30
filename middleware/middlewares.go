package middleware

import (
	"fmt"
	"net/http"
)

// func ChainMixed(args ...interface{}) http.Handler {
// 	var finalHandler http.Handler

// 	for i := len(args) - 1; i >= 0; i-- {
// 		arg := args[i]
// 		switch v := arg.(type) {
// 		case func(http.ResponseWriter, *http.Request):
// 			finalHandler = http.HandlerFunc(v)
// 		case func(http.Handler) http.Handler:
// 			if finalHandler == nil {
// 				panic("No handler found to apply middleware")
// 			}
// 			finalHandler = v(finalHandler)
// 		default:
// 			panic("Unknown type")
// 		}
// 	}

// 	return finalHandler
// }

func ExecutionTimeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("ExecutionTimeMiddleware start")
		next.ServeHTTP(w, r)
		fmt.Println("ExecutionTimeMiddleware end") // optional
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Log Middleware GET", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("AuthMiddleware executed")
		next.ServeHTTP(w, r)
	})
}
