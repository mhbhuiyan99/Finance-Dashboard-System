package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/mhbhuiyan99/Finance-Dashboard-System/config"
	"github.com/mhbhuiyan99/Finance-Dashboard-System/rest/handlers/record"
	"github.com/mhbhuiyan99/Finance-Dashboard-System/rest/handlers/user"

	middleware "github.com/mhbhuiyan99/Finance-Dashboard-System/rest/middlewares"
)

type Server struct {
	cnf *config.Config
	userHandler *user.Handler
	recordHandler *record.Handler
	httpServer *http.Server
}

func NewServer(
	cnf *config.Config,
	userHandler *user.Handler,
	recordHandler *record.Handler,
) *Server {
	return &Server{
		cnf: cnf,
		userHandler: userHandler,
		recordHandler: recordHandler,
	}
}

func (server *Server) Start() error {
	manager := middleware.NewManager()

	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	// initialize routes
	server.recordHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(server.cnf.HttpPort)
	fmt.Println("Server is running on port ", addr)

	server.httpServer = &http.Server{
		Addr:    addr,
		Handler: wrappedMux,
	}

	err := server.httpServer.ListenAndServe()
	return err
}