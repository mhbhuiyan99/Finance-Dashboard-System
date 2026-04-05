package user

import (
	"github.com/mhbhuiyan99/Finance-Dashboard-System/config"
	middleware "github.com/mhbhuiyan99/Finance-Dashboard-System/rest/middlewares"
)

type Handler struct {
	cnf *config.Config
	middlewares *middleware.Middlewares
	svc Service
}

func NewHandler(
	cnf *config.Config,
	middlewares *middleware.Middlewares,
	svc Service,
) *Handler {
	return &Handler{
		cnf: cnf,
		middlewares: middlewares,
		svc: svc,
	}	
}