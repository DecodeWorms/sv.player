package main

import (
	"fmt"
	"log"

	pr "github.com/DecodeWorms/messaging-protocol/pulse"
	"github.com/DecodeWorms/sv.player/config"
	db "github.com/DecodeWorms/sv.player/db/postgres"
	"github.com/DecodeWorms/sv.player/grpc"
)

func main() {
	cfg := config.ImportConfig(config.Config{})
	log.Printf("Starting the application with its dependencies at port : %s", cfg.ServerPort)
	db, err := db.New(cfg.DatabaseHost, cfg.DatabaseUserName, cfg.DatabaseName, cfg.DatabasePort, cfg.DatabasePassword)
	if err != nil {
		log.Println(err)
		return
	}
	pu, err := pr.NewMessage(cfg.PulsarUrl)
	if err != nil {
		log.Println(err)
		return
	}

	add := fmt.Sprintf(":%s", cfg.ServerPort)

	sr := grpc.NewServer(db, pu)
	err = sr.Run(add)
	log.Println(err)
}
