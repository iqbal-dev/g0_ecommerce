package middleware

import (
	"fmt"
	"net/http"
)

// Middleware defines the signature for HTTP middlewares
type Middleware func(http.Handler) http.Handler

// Manager holds the ServeMux and registered global middlewares
type Manager struct {
	mux               *http.ServeMux
	globalMiddlewares []Middleware
}

// NewManager creates a new Manager with a given ServeMux
func NewManager(mux *http.ServeMux) *Manager {
	return &Manager{
		mux:               mux,
		globalMiddlewares: []Middleware{},
	}
}

// Use registers one or more global middlewares
func (mgr *Manager) Use(mws ...Middleware) {
	mgr.globalMiddlewares = append(mgr.globalMiddlewares, mws...)
}

// chain wraps a handler with global and route-level middlewares
func (mgr *Manager) chain(final http.Handler, routeMws []Middleware) http.Handler {
	// Apply global middlewares first
	for _, mw := range mgr.globalMiddlewares {
		final = mw(final)
	}

	// Apply route-level middlewares
	for _, mw := range routeMws {
		final = mw(final)
	}

	return final
}

// GET registers a route for the GET HTTP method with optional middlewares
func (mgr *Manager) GET(path string, handlers ...interface{}) {
	if len(handlers) == 0 {
		panic("no handler provided for GET " + path)
	}

	var routeMws []Middleware
	var finalHandler http.Handler

	// Iterate from last to first to maintain declaration order
	for i := len(handlers) - 1; i >= 0; i-- {
		switch h := handlers[i].(type) {
		case Middleware:
			routeMws = append(routeMws, h)
		case func(http.ResponseWriter, *http.Request):
			finalHandler = http.HandlerFunc(h)
		case http.Handler:
			finalHandler = h
		default:
			panic(fmt.Sprintf("invalid type %T for GET %s", h, path))
		}
	}

	if finalHandler == nil {
		panic("no valid handler found for GET " + path)
	}

	// Wrap with middlewares
	finalHandler = mgr.chain(finalHandler, routeMws)

	// Restrict to GET method
	mgr.mux.Handle(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		finalHandler.ServeHTTP(w, r)
	}))
}

// ServeHTTP implements http.Handler to serve requests using the underlying mux
func (mgr *Manager) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mgr.mux.ServeHTTP(w, r)
}
