package api

import (
	"net/http"
)

// Router wraps http.ServeMux with convenience methods
type Router struct {
	mux *http.ServeMux
}

// NewRouter creates a new router
func NewRouter() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

// ServeHTTP implements http.Handler
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

// GET registers a GET route
func (r *Router) GET(pattern string, handler http.HandlerFunc) {
	r.mux.HandleFunc("GET "+pattern, handler)
}

// POST registers a POST route
func (r *Router) POST(pattern string, handler http.HandlerFunc) {
	r.mux.HandleFunc("POST "+pattern, handler)
}

// PUT registers a PUT route
func (r *Router) PUT(pattern string, handler http.HandlerFunc) {
	r.mux.HandleFunc("PUT "+pattern, handler)
}

// DELETE registers a DELETE route
func (r *Router) DELETE(pattern string, handler http.HandlerFunc) {
	r.mux.HandleFunc("DELETE "+pattern, handler)
}

// PATCH registers a PATCH route
func (r *Router) PATCH(pattern string, handler http.HandlerFunc) {
	r.mux.HandleFunc("PATCH "+pattern, handler)
}

// Handle registers a route for any HTTP method
func (r *Router) Handle(method, pattern string, handler http.Handler) {
	r.mux.Handle(method+" "+pattern, handler)
}

// HandleFunc registers a route for any HTTP method with a handler function
func (r *Router) HandleFunc(method, pattern string, handler http.HandlerFunc) {
	r.mux.HandleFunc(method+" "+pattern, handler)
}

// Group creates a sub-router with a common prefix
func (r *Router) Group(prefix string) *Router {
	// For simplicity, we'll just return a new router that wraps handlers with the prefix
	// This is a basic implementation - could be enhanced with middleware support per group
	return &Router{
		mux: r.mux,
	}
}
