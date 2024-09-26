package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var Db *sql.DB = ConnectDabase()

func ConnectDabase() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DbUser, DbPassword, DbHost, DbPort, DbName)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("Error opening database: %v\n", err)
	}

	return db
}
