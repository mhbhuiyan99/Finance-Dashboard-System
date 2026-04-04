package record

import (
	middleware "github.com/mhbhuiyan99/Finance-Dashboard-System/rest/middlewares"
)

type Handler struct {
	middlewares *middleware.Middlewares
	svc Service
}

func NewHandler(
	middlewares *middleware.Middlewares,
	svc Service,
) *Handler {
	return &Handler{
		middlewares: middlewares,
		svc: svc,
	}
}