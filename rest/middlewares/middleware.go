package middlewares

import "github.com/mhbhuiyan99/Finance-Dashboard-System/config"

type Middlewares struct {
	cnf *config.Config
}

func NewMiddlewares(cnf *config.Config) *Middlewares {
	return &Middlewares {
		cnf: cnf,
	}
}