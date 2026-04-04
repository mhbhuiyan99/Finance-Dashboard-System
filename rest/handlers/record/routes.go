package record

import (
	"net/http"

	"github.com/mhbhuiyan99/Finance-Dashboard-System/rest/middlewares"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {

	mux.Handle(
		"POST /records",
		manager.With(
			http.HandlerFunc(h.CreateRecord),
			h.middlewares.AuthenticateJWT,
			h.middlewares.RateLimit,
		),
	)	
		
}