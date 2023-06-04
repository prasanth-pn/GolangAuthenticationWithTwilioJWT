package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const (
	DBsource = "postgresql://jwt:1234@localhost:5432/jwtauth?sslmode=disable"
)

func ConnectDB() *sql.DB {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("errot connecting to database %s", err)
		}
	}()
	db, err := sql.Open("postgres", DBsource)

	if err != nil {
		panic(err)
	}
	return db

}
