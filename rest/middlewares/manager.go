package middlewares

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: []Middleware{},
	}
}

func (m *Manager) Use(middlewares ...Middleware) {
	m.globalMiddlewares = append(m.globalMiddlewares, middlewares...)
}

// Taking multiple middlewares and chaining them together
// With — applies to specific routes only
func (m *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	handler := next
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

// WrapMux — applies to ALL routes globally
func (m *Manager) WrapMux(handler http.Handler) http.Handler {
	wrappedHandler := handler
	for _, middleware := range m.globalMiddlewares {
		wrappedHandler = middleware(wrappedHandler)
	}
	return wrappedHandler
}