package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "shop_db2"
)

func InitDB() *sql.DB{
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s application_name=demo_practice sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("pgx", psqlInfo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable to ping:%v\n", err)
	}
	return db
}