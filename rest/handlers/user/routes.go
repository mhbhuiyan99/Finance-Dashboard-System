package user

import (
	"net/http"

	"github.com/mhbhuiyan99/Finance-Dashboard-System/rest/middlewares"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {

	mux.Handle(
		"POST /users",
		manager.With(
			http.HandlerFunc(h.CreateUser),
			h.middlewares.RateLimit,
		),
	)	
		
}