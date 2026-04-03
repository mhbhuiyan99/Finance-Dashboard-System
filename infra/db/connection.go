package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mhbhuiyan99/Finance-Dashboard-System/config"
)

func GetConnectionString(cnf *config.DBConfig) string {

	connString := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%d",
		cnf.User,
		cnf.Password,
		cnf.Name,
		cnf.Host,
		cnf.Port,
	)

	if !cnf.EnableSSLMode {
		connString += " sslmode=disable"
	}

	return connString
}

func NewDBConnection(cnf *config.DBConfig) (*sqlx.DB, error) {
	dbSource := GetConnectionString(cnf)

	db, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}