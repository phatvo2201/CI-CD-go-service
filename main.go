package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/phatvo2201/simplebank/api"
	db "github.com/phatvo2201/simplebank/db/sqlc"
	"github.com/phatvo2201/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	server.Start(config.SVAddress)
}
