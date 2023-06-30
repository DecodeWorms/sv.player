package main

import (
	"log"

	store "github.com/DecodeWorms/messaging-protocol"
	pr "github.com/DecodeWorms/messaging-protocol/pulse"
	"github.com/DecodeWorms/sv.player/config"
	db "github.com/DecodeWorms/sv.player/db/postgres"
	"github.com/DecodeWorms/sv.player/server"
)

func main() {
	cfg := config.ImportConfig(config.Config{})
	serv, err := db.New(cfg.DatabaseHost, cfg.DatabaseUserName, cfg.DatabaseName, cfg.DatabasePort, cfg.DatabasePassword)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = pr.Init(store.Options{
		Address: cfg.PulsarUrl,
	})
	if err != nil {
		log.Println(err)
		return
	}

	h, err := server.New(*serv)
	if err != nil {
		log.Println(err)
		return
	}
	//migrate the database tables to the Postgres Server
	if err = h.CreateTableMigration(); err != nil {
		log.Println("error migrating database tables")
	}

}
