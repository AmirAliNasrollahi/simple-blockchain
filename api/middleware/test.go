package middleware

import (
	"fmt"
	"net/http"
)


func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "amirali" {
			fmt.Fprintf(w, "your token (%v) is not valid", r.Header.Get("Authorization"))
			return
		}
		next.ServeHTTP(w, r)
	})	
}
