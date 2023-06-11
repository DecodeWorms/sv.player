package server

import (
	messaging "github.com/DecodeWorms/messaging-protocol"
	"github.com/DecodeWorms/sv.player/db/postgres"
)

type PlayerHandler struct {
	playerService postgres.PostgresStore
	eventStore    messaging.PulsarStore
	//logger logger.Logger
}

func New(p postgres.PostgresStore, store messaging.PulsarStore) *PlayerHandler {
	return &PlayerHandler{
		playerService: p,
		eventStore:    store,
	}
}

func (p PlayerHandler) CreateTableMigration() error {
	//create table for players personal info
	err := p.playerService.AutoMigratePersonalInfo()
	if err != nil {
		return err
	}
	if err = p.playerService.AutoMigrateFieldInfo(); err != nil {
		return err
	}
	if err = p.playerService.AutoMigrateAddressInfo(); err != nil {
		return err
	}
	if err = p.playerService.AutoMigrateClubsPreviouslyPlazed(); err != nil {
		return err
	}
	return nil
}
