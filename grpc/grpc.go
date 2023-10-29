package grpc

import (
	"log"
	"net"

	"github.com/DecodeWorms/messaging-protocol/pulse"
	"github.com/DecodeWorms/sv.player/config"
	"github.com/DecodeWorms/sv.player/db"
	"github.com/DecodeWorms/sv.player/pb/protos/pb/player"
	"github.com/DecodeWorms/sv.player/server"
	"google.golang.org/grpc"
)

type Server struct {
	server *grpc.Server
	player.UnimplementedPlayerServiceServer
}

func NewServer(db db.DataStore, pulStore *pulse.Message) *Server {
	ser := grpc.NewServer()
	handler, err := server.NewPlayerHandler(db, pulStore)
	if err != nil {
		log.Println(err)
	}
	err = handler.CreateTableMigration()
	if err != nil {
		log.Println()
	}
	//register player service
	player.RegisterPlayerServiceServer(ser, handler)

	return &Server{
		server: ser,
	}

}

func (s Server) Run(addr string) error {
	cfg := config.ImportConfig(config.Config{})
	listen, err := net.Listen(cfg.ServerProtocol, addr)
	if err != nil {
		return err
	}
	return s.server.Serve(listen)
}
