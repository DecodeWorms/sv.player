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
	Dbname   = "soccermetrics"
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
		log.Println("unabe to create a table player personal info")
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
