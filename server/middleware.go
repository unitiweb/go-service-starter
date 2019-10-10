package server

import (
	"net/http"
)

// The structure of a Middleware
type Middleware struct {
	Handler func(w http.ResponseWriter, r *http.Request, next http.Handler)
}

// Holds all the configured middleware
var middleware []Middleware

// Add middleware to the server
func AddMiddleware(h func(w http.ResponseWriter, r *http.Request, next http.Handler)) {
	middleware = append(middleware, Middleware{Handler: h})
}

// Middleware to add the proper response headers
func responseMiddleware(w http.ResponseWriter, r *http.Request, next http.Handler) {
	w.Header().Set("Content-Type", "application/json")
	next.ServeHTTP(w, r)
}

// Add all configured middleware to the server
func loadMiddleware() {
	// Load al the configured middleware
	for _, m := range middleware {
		router.Use(
			func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					m.Handler(w, r, next)
				})
			},
		)
	}
}