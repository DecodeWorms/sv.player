package db

import (
	"log"
	"testing"

	"github.com/DecodeWorms/server-contract/models"
	"github.com/stretchr/testify/assert"
)

const (
	Host     = "localhost"
	User     = "runner" //user = "abdulhmeed"
	Password = "password"
	Dbname   = "services" //soccermetrics
	port     = "5432"
)

func TestUpdatePlayer(t *testing.T) {
	//Establish connection  to the Db..
	db, err := New(Host, User, Dbname, port, Password)
	if err != nil {
		log.Println("Unable to connect to the db..")
	}

	//player id
	id := "Qos-1234-test"
	playerRecord := &models.PersonalInfo{
		FirstName:   "Abdulhameed",
		LastName:    "Awwal",
		Gender:      "male",
		PhoneNumber: "0919673928430",
	}

	err = db.UpdatePlayer(id, playerRecord)
	assert.Nil(t, err)
}

func TestGetPlayerById(t *testing.T) {
	//Establish connection  to the Db..
	db, err := New(Host, User, Dbname, port, Password)
	if err != nil {
		log.Println("Unable to connect to the db..")
	}

	//user id
	playerId := "Qos-1234-test"

	user, err := db.GetPlayerById(playerId)
	assert.Nil(t, err)
	assert.NotNil(t, user)

}

func TestGetPlayerByPhoneNumber(t *testing.T) {
	//Establish connection  to the Db..
	db, err := New(Host, User, Dbname, port, Password)
	if err != nil {
		log.Println("Unable to connect to the db..")
	}

	//user id
	phoneNumber := "0919673928430"

	u, err := db.GetPlayerByPhoneNumber(phoneNumber)
	assert.Nil(t, err)
	assert.NotNil(t, u)

}

func TestCreatePlayerWithField(t *testing.T) {
	//Establish connection  to the Db..
	db, err := New(Host, User, Dbname, port, Password)
	if err != nil {
		log.Println("Unable to connect to the db..")
	}
	//create a database table for a player field record on Github action
	if err = db.db.AutoMigrate(&models.FieldInfo{}); err != nil {
		log.Println("Unabe to create a table player personal info")
	}
	playerRecord := &models.FieldInfo{
		PersonalInfoId:      "user-id-1234-fati",
		YearOfExperience:    "6 years",
		NumberOfGoalsScored: "50",
	}

	err = db.CreatePlayerWithFieldsData(*playerRecord)
	assert.Nil(t, err)

}

func TestUpdateFieldRecord(t *testing.T) {
	//Establish connection  to the Db..
	db, err := New(Host, User, Dbname, port, Password)
	if err != nil {
		log.Println("Unable to connect to the db..")
	}
	playerRecord := &models.FieldInfo{
		YearOfExperience:    "6 years",
		NumberOfGoalsScored: "50",
		JerseyNumber:        "22",
		YearJoined:          "2018-12-01",
	}
	err = db.UpdatePlayerWithFieldsInfo("user-id-1234-fati", playerRecord)
	assert.Nil(t, err)
}

func TestGetPlayerWithFieldsInfoById(t *testing.T) {
	//Establish connection  to the Db..
	db, err := New(Host, User, Dbname, port, Password)
	if err != nil {
		log.Println("Unable to connect to the db..")
	}
	userId := "user-id-1234-fati"
	res, err := db.GetPlayerWithFieldsInfoById(userId)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDeletePlayer(t *testing.T) {
	//Establish connection  to the Db..
	db, err := New(Host, User, Dbname, port, Password)
	if err != nil {
		log.Println("Unable to connect to the db..")
	}
	userId := "user-id-1234-fati"
	err = db.DeletePlayer(userId)
	assert.Nil(t, err)

}
