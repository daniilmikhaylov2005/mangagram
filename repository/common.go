package repository

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func getConnection() *sql.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}
	db, err := sql.Open("postgres", os.Getenv("DB"))
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
