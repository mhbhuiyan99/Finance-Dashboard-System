package cmd

import (
	"fmt"
	"os"

	"github.com/mhbhuiyan99/Finance-Dashboard-System/config"
	"github.com/mhbhuiyan99/Finance-Dashboard-System/infra/db"
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

	fmt.Println("Server is running...")
}