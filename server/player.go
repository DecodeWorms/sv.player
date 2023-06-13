package server

import (
	pr "github.com/DecodeWorms/messaging-protocol"
	db "github.com/DecodeWorms/sv.player/db/postgres"
)

type PlayerHandler struct {
	playerService db.PostgresStore
	store         pr.PulsarStore
	//add the logger
}

func New(p db.PostgresStore) (PlayerHandler, error) {
	return PlayerHandler{
		playerService: p,
	}, nil
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
