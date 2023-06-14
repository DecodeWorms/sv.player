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

func TestCreatePlayer(t *testing.T) {
	//Establish connection  to the Db..
	db, err := New(Host, User, Dbname, port, Password)
	if err != nil {
		log.Println("Unable to connect to the db..")
	}
	//create a database table for a player on Github action
	if err = db.db.AutoMigrate(&models.PersonalInfo{}); err != nil {
		log.Println("Unabe to create a table player personal info")
	}

	//Persist Data to the db..
	playerRecord := models.PersonalInfo{
		Id:        "user-id-1234-dan",
		FirstName: "Danny",
		LastName:  "Ryan",
		Gender:    "male",
	}
	var p = PostgresStore{
		db: db.db,
	}
	if err = p.CreatePlayer(playerRecord); err != nil {
		assert.Nil(t, err)
	}
}

func TestGetPlayerById(t *testing.T) {
	//Establish connection  to the Db..
	db, err := New(Host, User, Dbname, port, Password)
	if err != nil {
		log.Println("Unable to connect to the db..")
	}
	//create a database table for a player on Github action
	if err = db.db.AutoMigrate(&models.PersonalInfo{}); err != nil {
		log.Println("Unabe to create a table player personal info")
	}

	//user id
	playerId := "user-id-1234-dan"

	var p = PostgresStore{
		db: db.db,
	}
	user, err := p.GetPlayerById(playerId)
	assert.Nil(t, err)
	assert.NotNil(t, user)

}

func TestGetPlayerByPhoneNumber(t *testing.T) {
	//Establish connection  to the Db..
	db, err := New(Host, User, Dbname, port, Password)
	if err != nil {
		log.Println("Unable to connect to the db..")
	}
	//create a database table for a player on Github action
	if err = db.db.AutoMigrate(&models.PersonalInfo{}); err != nil {
		log.Println("unabe to create a table player personal info")
	}

	//user id
	phoneNumber := "09000000000"

	playerRecord := models.PersonalInfo{
		Id:          "user-id-1234-fun",
		FirstName:   "Funmi",
		LastName:    "Adeola",
		Gender:      "female",
		PhoneNumber: "09000000000",
	}
	var p = PostgresStore{
		db: db.db,
	}
	err = p.CreatePlayer(playerRecord)
	assert.Nil(t, err)
	u, err := p.GetPlayerByPhoneNumber(phoneNumber)
	assert.Nil(t, err)
	assert.Equal(t, u.FirstName, playerRecord.FirstName)
	assert.Equal(t, u.LastName, playerRecord.LastName)

}

func TestUpdatePlayer(t *testing.T) {
	//Establish connection  to the Db..
	db, err := New(Host, User, Dbname, port, Password)
	if err != nil {
		log.Println("Unable to connect to the db..")
	}
	//create a database table for a player on Github action
	if err = db.db.AutoMigrate(&models.PersonalInfo{}); err != nil {
		log.Println("Unabe to create a table player personal info")
	}

	//player id
	id := "user-id-1234-fun"
	playerRecord := &models.PersonalInfo{
		FirstName:   "Funmilayo",
		LastName:    "Akinlola",
		Gender:      "female",
		PhoneNumber: "09000000000",
	}
	var p = PostgresStore{
		db: db.db,
	}
	err = p.UpdatePlayer(id, playerRecord)
	assert.Nil(t, err)
}

/*func TestCreatePlayerWithField(t *testing.T) {
	//Establish connection  to the Db..
	db, err := New(Host, User, Dbname, port, Password)
	if err != nil {
		log.Println("Unable to connect to the db..")
	}
	//create a database table for a player on Github action
	if err = db.db.AutoMigrate(&models.FieldInfo{}); err != nil {
		log.Println("unabe to create a table player personal info")
	}
	playerRecord := &models.FieldInfo{
		Id:                  "test-player-id-funmilayo",
		PersonalInfoId:      "user-id-1234-fun",
		YearOfExperience:    "6 years",
		NumberOfGoalsScored: 50000,
	}
	var p = PostgresStore{
		db: db.db,
	}
	err = p.CreatePlayerWithFieldsData(*playerRecord)
	assert.Nil(t, err)

}
*/
