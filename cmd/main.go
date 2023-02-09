package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/zhas-off/simple-bank-service/api"
	db "github.com/zhas-off/simple-bank-service/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:1234@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8085"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
