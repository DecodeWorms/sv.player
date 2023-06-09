package db

import (
	"fmt"
	"log"
	"time"

	"github.com/DecodeWorms/server-contract/models"
	"github.com/DecodeWorms/sv.player/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _ db.DataStore = &PostgresStore{}

type PostgresStore struct {
	db *gorm.DB
	//add the logger
}

func New(host, user, name, port, password string) (*PostgresStore, error) {
	log.Println("Connecting to the DB...")

	uri := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", host, user, name, port, password)
	database, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Connected to the Db..")
	return &PostgresStore{
		db: database,
	}, nil
}

func (p PostgresStore) AutoMigratePersonalInfo() error {
	per := models.PersonalInfo{}
	err := p.db.AutoMigrate(&per)
	return err
}

func (p PostgresStore) AutoMigrateFieldInfo() error {
	field := models.FieldInfo{}
	err := p.db.AutoMigrate(&field)
	return err
}

func (p PostgresStore) AutoMigrateAddressInfo() error {
	add := models.Address{}
	err := p.db.AutoMigrate(&add)
	return err
}

func (p PostgresStore) AutoMigrateClubsPreviouslyPlazed() error {
	clubs := models.ClubsHePreviouslyPlayed{}
	err := p.db.AutoMigrate(&clubs)
	return err
}
func (p PostgresStore) CreatePlayer(data models.PersonalInfo) error {
	data.CreatedAt = time.Now()
	err := p.db.Create(&models.PersonalInfo{
		Id:            data.Id,
		FirstName:     data.FirstName,
		LastName:      data.LastName,
		Gender:        data.Gender,
		MaritalStatus: data.MaritalStatus,
		Email:         data.Email,
		PhoneNumber:   data.PhoneNumber,
	}).Error
	return err
}

func (p PostgresStore) GetPlayerById(id string) (*models.PersonalInfo, error) {
	player := models.PersonalInfo{}
	err := p.db.Where("id = ?", id).First(&player).Error
	return &player, err
}

func (p PostgresStore) GetPlayerByPhoneNumber(phoneNumber string) (*models.PersonalInfo, error) {
	player := models.PersonalInfo{}
	err := p.db.Where("phone_number = ?", phoneNumber).First(&player).Error
	return &player, err
}

func (p PostgresStore) UpdatePlayer(id string, data *models.PersonalInfo) error {
	data.UpdatedAt = time.Now()
	var old = &models.PersonalInfo{}
	_ = p.db.Where("id = ?", id).First(old).Error
	d := buildPlayerPayload(old, data)
	err := p.db.Model(&models.PersonalInfo{}).Where("id = ?", id).Updates(d).Error
	return err
}

func (p PostgresStore) CreatePlayerWithFieldsData(data models.FieldInfo) error {
	data.CreatedAt = time.Now()
	err := p.db.Create(&models.FieldInfo{
		PersonalInfoId:      data.PersonalInfoId,
		YearOfExperience:    data.YearOfExperience,
		NumberOfGoalsScored: data.NumberOfGoalsScored,
		JerseyNumber:        data.JerseyNumber,
		YearJoined:          data.YearJoined,
		PositionOnTheField:  data.PositionOnTheField,
		PlayerClubStatus:    data.PlayerClubStatus,
	}).Error
	return err
}

func (p PostgresStore) GetPlayer(jerseyNumber string) (*models.FieldInfo, error) {
	var player = &models.FieldInfo{}
	err := p.db.Where("jersey_number = ?", jerseyNumber).First(player).Error
	return player, err

}

func (p PostgresStore) UpdatePlayerWithFieldsInfo(id string, data *models.FieldInfo) error {
	data.UpdatedAt = time.Now()

	var old = &models.FieldInfo{}
	_ = p.db.Where("personal_info_id = ?", id).First(old).Error
	d := buildPlayerWithFieldPayload(old, data)
	err := p.db.Model(&models.FieldInfo{}).Where("personal_info_id = ?", id).Updates(d).Error
	return err
}
func (p PostgresStore) GetPlayerWithFieldsInfoById(id string) (*models.FieldInfo, error) {
	var playerWithField = &models.FieldInfo{}
	err := p.db.Where("personal_info_id = ?", id).First(playerWithField).Error
	return playerWithField, err

}

func (p PostgresStore) DeletePlayer(id string) error {
	player := models.FieldInfo{}
	err := p.db.Where("personal_info_id = ?", id).Find(&player).Delete(&models.FieldInfo{}).Error
	return err
}

func buildPlayerPayload(old, new *models.PersonalInfo) *models.PersonalInfo {
	if new == nil {
		return nil
	}

	if new.FirstName != "" {
		old.FirstName = new.FirstName
	}
	if new.LastName != "" {
		old.LastName = new.FirstName
	}
	if new.Gender != "" {
		old.Gender = new.Gender
	}
	if new.MaritalStatus != "" {
		old.MaritalStatus = new.MaritalStatus
	}
	if new.Email != "" {
		old.Email = new.Email
	}
	if new.Address.Name != "" {
		old.Address.Name = new.Address.Name
	}
	if new.Address.City != "" {
		old.Address.City = new.Address.City
	}
	if old.Address.ZipCode != "" {
		old.Address.ZipCode = new.Address.ZipCode
	}
	return old
}

func buildPlayerWithFieldPayload(old, new *models.FieldInfo) *models.FieldInfo {
	if new == nil {
		return nil
	}
	if new.JerseyNumber != 0 {
		old.JerseyNumber = new.JerseyNumber
	}
	if new.NumberOfGoalsScored != 0 {
		old.NumberOfGoalsScored = new.NumberOfGoalsScored
	}
	if new.PlayerClubStatus != "" {
		old.PlayerClubStatus = new.PlayerClubStatus
	}
	if new.PositionOnTheField != "" {
		old.PositionOnTheField = new.PositionOnTheField
	}
	if new.YearJoined != "" {
		old.YearJoined = new.YearJoined
	}
	if new.YearOfExperience != "" {
		old.YearOfExperience = new.YearOfExperience
	}
	return old
}
