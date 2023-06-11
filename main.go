package main

import (
	"log"

	store "github.com/DecodeWorms/messaging-protocol"
	"github.com/DecodeWorms/messaging-protocol/pulse"
	"github.com/DecodeWorms/sv.player/config"
	"github.com/DecodeWorms/sv.player/db/postgres"
	"github.com/DecodeWorms/sv.player/server"
)

func main() {
	cfg := config.ImportConfig(config.Config{})
	pStore, err := postgres.New(cfg.DatabaseHost, cfg.DatabaseUserName, cfg.DatabaseName, cfg.DatabasePort)
	if err != nil {
		log.Println(err)
	}
	p, err := pulse.Init(store.Options{
		Address: cfg.PulsarUrl,
	})
	if err != nil {
		log.Println(err)
	}
	h := server.New(*pStore, p)
	if err = h.CreateTableMigration(); err != nil {
		log.Printf("error creating tables %v", err)
		return
	}
	log.Println("Tables created successfully")

}
