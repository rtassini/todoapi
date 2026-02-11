package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "go_db"
)

func Connect() (*sql.DB, error) {
	//host := os.Getenv("DB_HOST")
	//portStr := os.Getenv("DB_PORT")
	//user := os.Getenv("DB_USER")
	//password := os.Getenv("DB_PASSWORD")
	//dbname := os.Getenv("DB_NAME")
	//
	//port, err := strconv.Atoi(portStr)
	//if err != nil {
	//	panic(err)
	//}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Error pinging database: %v", err)
		panic(err)
	}
	return db, nil
}
