package routes

import (
	"fmt"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type routeEntry struct {
	method  string
	handler http.Handler
}

type Manager struct {
	mux               *http.ServeMux
	globalMiddlewares []Middleware
	routes            map[string][]routeEntry // key = path
}

func NewManager(mux *http.ServeMux) *Manager {
	return &Manager{
		mux:               mux,
		globalMiddlewares: []Middleware{},
		routes:            make(map[string][]routeEntry),
	}
}

// chain wraps a handler with global and route-level middlewares
func (mgr *Manager) chain(final http.Handler, routeMws []Middleware) http.Handler {
	for _, mw := range mgr.globalMiddlewares {
		final = mw(final)
	}
	for _, mw := range routeMws {
		final = mw(final)
	}
	return final
}

func (mgr *Manager) registerRoute(method, path string, handlers ...interface{}) {
	if len(handlers) == 0 {
		panic(fmt.Sprintf("no handler provided for %s %s", method, path))
	}

	var routeMws []Middleware
	var finalHandler http.Handler
	for i := len(handlers) - 1; i >= 0; i-- {
		switch h := handlers[i].(type) {
		case func(http.Handler) http.Handler:
			routeMws = append(routeMws, h)
		case func(http.ResponseWriter, *http.Request):
			finalHandler = http.HandlerFunc(h)
		case http.Handler:
			finalHandler = h
		default:
			panic(fmt.Sprintf("invalid type %T for GET %s", h, path))
		}
	}

	finalHandler = mgr.chain(finalHandler, routeMws)

	// Store handler per method
	mgr.routes[path] = append(mgr.routes[path], routeEntry{method: method, handler: finalHandler})

	// Register a single mux handler per path if not registered yet
	if len(mgr.routes[path]) == 1 {
		mgr.mux.Handle(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			entries := mgr.routes[path]
			for _, e := range entries {
				if r.Method == e.method {
					e.handler.ServeHTTP(w, r)
					return
				}
			}
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}))
	}
}

// Public methods for all HTTP verbs
func (mgr *Manager) GET(path string, handlers ...interface{}) {
	mgr.registerRoute(http.MethodGet, path, handlers...)
}
func (mgr *Manager) POST(path string, handlers ...interface{}) {
	mgr.registerRoute(http.MethodPost, path, handlers...)
}
func (mgr *Manager) PUT(path string, handlers ...interface{}) {
	mgr.registerRoute(http.MethodPut, path, handlers...)
}
func (mgr *Manager) PATCH(path string, handlers ...interface{}) {
	mgr.registerRoute(http.MethodPatch, path, handlers...)
}
func (mgr *Manager) DELETE(path string, handlers ...interface{}) {
	mgr.registerRoute(http.MethodDelete, path, handlers...)
} // ServeHTTP implements http.Handler
func (mgr *Manager) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mgr.mux.ServeHTTP(w, r)
}
