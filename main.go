package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/saurabhsingh121/simplebank/api"
	db "github.com/saurabhsingh121/simplebank/db/sqlc"
	"github.com/saurabhsingh121/simplebank/util"
)


func main(){
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configuration: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start the server: ", err)
	}
}