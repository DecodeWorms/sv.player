package postgres

import (
	"fmt"
	"log"

	"github.com/DecodeWorms/sv.player/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _ db.DataStore = &PostgresStore{}

type PostgresStore struct {
	db *gorm.DB
	//add the logger
	//add the
}

func New(host, user, name, port string) (*PostgresStore, error) {
	log.Println("Connecting to the DB...")
	uri := fmt.Sprintf("host=%s user=%s dbname=%s port=%s", host, user, name, port)
	database, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("Connected to the Db..")
	return &PostgresStore{
		db: database,
	}, nil
}
