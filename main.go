package main

import (
	"log"

	"github.com/DecodeWorms/sv.player/config"
	"github.com/DecodeWorms/sv.player/db/postgres"
)

func main() {
	cfg := config.ImportConfig(config.Config{})
	_, err := postgres.New(cfg.DatabaseHost, cfg.DatabaseUserName, cfg.DatabaseName, cfg.DatabasePort)
	if err != nil {
		log.Println(err)
	}

}
