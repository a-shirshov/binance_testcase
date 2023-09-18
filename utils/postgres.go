package utils

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitPostgres() (*sqlx.DB, error) {
	user := "artyom"
	database := "postgres"
	password := "12345678"
	host := "postgres"
	port := 5432

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, database)

	
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return sqlx.NewDb(db, "postgres"), nil
}