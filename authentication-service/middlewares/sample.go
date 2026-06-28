package middlewares

import (
	"fmt"
	"net/http"
)
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		fmt.Println("Receive Request", r.Method,r.URL.Path)
		next.ServeHTTP(w,r)
	})
}