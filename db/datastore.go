package db

import "github.com/DecodeWorms/server-contract/models"

// DataStore interface
//
//go:generate mockgen -source=datastore.go -destination=../mocks/datastore_mock.go -package=mocks
type DataStore interface {
	PlayerStore
}
type PlayerStore interface {
	AutoMigratePersonalInfo() error
	AutoMigrateFieldInfo() error
	AutoMigrateAddressInfo() error
	AutoMigrateClubsPreviouslyPlazed() error
	CreatePlayer(data models.PersonalInfo) error
	GetPlayerById(id string) (*models.PersonalInfo, error)
	GetPlayerByPhoneNumber(phoneNumber string) (*models.PersonalInfo, error)
	UpdatePlayer(id string, data *models.PersonalInfo) error
	CreatePlayerWithFieldsData(data models.FieldInfo) error
	GetPlayer(jerseyNumber string) (*models.FieldInfo, error)
	UpdatePlayerWithFieldsInfo(id string, data *models.FieldInfo) error
	GetPlayerWithFieldsInfoById(id string) (*models.FieldInfo, error)
	DeletePlayer(id string) error
	CreateAddress(data *models.Address) error
	UpdateAddress(id string, data *models.Address) error
	GetPlayerByJerseyNumber(jerseyNumber string) (*models.FieldInfo, error)
	DeletePlayerFieldInfo(id string) error
	DeletePlayerAddress(id string) error
	GetAddressById(id string) (*models.Address, error)
	GetPlayerByEmail(email string)(*models.PersonalInfo, error)
}
