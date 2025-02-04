package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/nizaliyev7/employess/api"
	db "github.com/nizaliyev7/employess/db/sqlc"
	"github.com/nizaliyev7/employess/db/util"
)

func main() {

	cfg, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(cfg.DBDriver, cfg.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	s := db.NewStore(conn)

	server := api.NewServer(s)

	err = server.Start(cfg.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
