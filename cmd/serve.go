package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/mhbhuiyan99/Finance-Dashboard-System/config"
	"github.com/mhbhuiyan99/Finance-Dashboard-System/infra/db"
	"github.com/mhbhuiyan99/Finance-Dashboard-System/rest"

	recordD "github.com/mhbhuiyan99/Finance-Dashboard-System/record"
	userD "github.com/mhbhuiyan99/Finance-Dashboard-System/user"

	"github.com/mhbhuiyan99/Finance-Dashboard-System/repo"
	"github.com/mhbhuiyan99/Finance-Dashboard-System/rest/handlers/record"
	"github.com/mhbhuiyan99/Finance-Dashboard-System/rest/handlers/user"
	middleware "github.com/mhbhuiyan99/Finance-Dashboard-System/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewDBConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	middlewares := middleware.NewMiddlewares(cnf)

	userRepo := repo.NewUserRepo(dbCon)
	recordRepo := repo.NewRecordRepo(dbCon)

	userSvc := userD.NewService(userRepo)
	rcdSvc := recordD.NewService(recordRepo)

	userHandler := user.NewHandler(cnf, middlewares, userSvc)
	recordHandler := record.NewHandler(middlewares, rcdSvc)
	

	server := rest.NewServer(
		cnf,
		userHandler,
		recordHandler,
	)

	if err := server.Start(); err != nil && err != http.ErrServerClosed {
		fmt.Println("Error starting server: ", err)
	}

	if err := dbCon.Close(); err != nil {
		fmt.Println("Error closing database connection: ", err)
	}
}