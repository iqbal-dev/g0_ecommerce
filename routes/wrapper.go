package routes

import "net/http"

// RouterWrapper wraps a ServeMux with a base prefix
type RouterWrapper struct {
	Router *http.ServeMux
	Prefix string
}

// Handle automatically prepends the prefix to the path
func (r *RouterWrapper) Handle(path string, handler http.Handler) {
	r.Router.Handle(r.Prefix+path, handler)
}

// HandleFunc convenience method
func (r *RouterWrapper) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) {
	r.Handle(path, http.HandlerFunc(f))
}
