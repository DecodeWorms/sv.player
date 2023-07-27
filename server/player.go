package server

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	//pr "github.com/DecodeWorms/messaging-protocol"
	"github.com/DecodeWorms/server-contract/models"
	data "github.com/DecodeWorms/sv.player/db"
	"github.com/DecodeWorms/sv.player/pb/protos/pb/player"
)

type PlayerHandler struct {
	playerService data.PlayerStore
	player.UnimplementedPlayerServiceServer
}

func NewPlayerHandler(p data.PlayerStore) (PlayerHandler, error) {
	return PlayerHandler{
		playerService: p,
	}, nil
}

var _ player.PlayerServiceServer = PlayerHandler{}

func (p PlayerHandler) CreatePlayer(ctx context.Context, in *player.CreatePlayerRequest) (*player.Empty, error) {
	//verify if the email already in use
	_, err := p.playerService.GetPlayerByEmail(in.Email)
	if err == nil {
		return nil, fmt.Errorf("error email already exist %v", nil)
	}

	//verify if the phone number already in use
	_, err = p.playerService.GetPlayerByPhoneNumber(in.PhoneNumber)
	if err == nil {
		return nil, fmt.Errorf("error user phone number already exist %v", nil)
	}
	data := models.PersonalInfo{
		Id:            generatePlayerId(11),
		FirstName:     in.FirstName,
		LastName:      in.LastName,
		Gender:        in.Gender,
		MaritalStatus: in.MaritalStatus,
		Email:         in.Email,
		PhoneNumber:   in.PhoneNumber,
	}
	if err := p.playerService.CreatePlayer(data); err != nil {
		return nil, fmt.Errorf("error creating a player personal info %v", err)
	}
	return &player.Empty{}, nil
}

func (p PlayerHandler) GetPlayerById(ctx context.Context, in *player.GetPlayerByIdRequest) (*player.GetPlayerByIdResponse, error) {
	//get player personal info
	playerId := in.Id
	res, err := p.playerService.GetPlayerById(playerId)
	if err != nil {
		return nil, fmt.Errorf("error getting player info %v", err)
	}

	//get player field info

	field, err := p.playerService.GetPlayerWithFieldsInfoById(playerId)
	if err != nil {
		return nil, fmt.Errorf("error getting player field info %v", err)
	}

	//get player address info

	add, err := p.playerService.GetAddressById(playerId)
	if err != nil {
		return nil, fmt.Errorf("error getting player address info %v", err)
	}

	return &player.GetPlayerByIdResponse{
		FirstName:           res.FirstName,
		LastName:            res.LastName,
		Gender:              res.Gender,
		PhoneNumber:         res.PhoneNumber,
		MaritalStatus:       res.MaritalStatus,
		Email:               res.Email,
		YearOfExperience:    field.YearOfExperience,
		NumberOfGoalsScored: field.NumberOfGoalsScored,
		JerseyNumber:        field.JerseyNumber,
		YearJoined:          field.YearJoined,
		PositionOnTheField:  field.PositionOnTheField,
		PlayerStatus:        field.PlayerClubStatus,
		Name:                add.Name,
		City:                add.City,
		ZipCode:             add.ZipCode,
	}, nil
}

func (p PlayerHandler) GetPlayerByPhoneNumber(ctx context.Context, in *player.GetPlayerByPhoneNumberRequest) (*player.GetPlayerByIdResponse, error) {
	playerId := in.PhoneNumber
	res, err := p.playerService.GetPlayerByPhoneNumber(playerId)
	if err != nil {
		return nil, fmt.Errorf("error getting player data %v", err)
	}
	field, err := p.playerService.GetPlayerWithFieldsInfoById(res.Id)
	if err != nil {
		return nil, fmt.Errorf("error getting player field info %v", err)
	}

	//get player address info
	add, err := p.playerService.GetAddressById(res.Id)
	if err != nil {
		return nil, fmt.Errorf("error getting player address info %v", err)
	}

	return &player.GetPlayerByIdResponse{
		FirstName:           res.FirstName,
		LastName:            res.LastName,
		Gender:              res.Gender,
		PhoneNumber:         res.PhoneNumber,
		MaritalStatus:       res.MaritalStatus,
		Email:               res.Email,
		YearOfExperience:    field.YearOfExperience,
		NumberOfGoalsScored: field.NumberOfGoalsScored,
		JerseyNumber:        field.JerseyNumber,
		YearJoined:          field.YearJoined,
		PositionOnTheField:  field.PositionOnTheField,
		PlayerStatus:        field.PlayerClubStatus,
		Name:                add.Name,
		City:                add.City,
		ZipCode:             add.ZipCode,
	}, nil
}

// complete me
func (p PlayerHandler) CompleteKyc(ctx context.Context, in *player.UpdateKyc) (*player.Empty, error) {
	//Verify if the player Id exists
	_, err := p.playerService.GetPlayerById(in.Id)
	if err != nil {
		return nil, errors.New("user record is not exist")
	}

	//Update field info
	data := models.FieldInfo{
		PersonalInfoId:      in.Id,
		YearOfExperience:    in.YearOfExperience,
		NumberOfGoalsScored: in.NumberOfGoalsScored,
		JerseyNumber:        in.JerseyNumber,
		YearJoined:          in.YearJoined,
		PositionOnTheField:  in.PositionOnTheField,
		PlayerClubStatus:    in.PlayerStatus,
	}

	if err := p.playerService.CreatePlayerWithFieldsData(data); err != nil {
		return nil, errors.New("unable to update player field record")
	}

	//update player address record
	add := &models.Address{
		PersonalInfoId: in.Id,
		Name:           in.Name,
		ZipCode:        in.ZipCode,
		City:           in.City,
	}
	if err := p.playerService.CreateAddress(add); err != nil {
		return nil, errors.New("unable to update player address")
	}
	return &player.Empty{}, nil
}

func (p PlayerHandler) UpdatePlayerExistingRecord(ctx context.Context, in *player.UpdatePlayerRequest) (*player.Empty, error) {
	//verify if the user exists
	_, err := p.playerService.GetPlayerById(in.Id)
	if err != nil {
		return nil, errors.New("error encountered in getting personal record")
	}

	//update personal info
	per := &models.PersonalInfo{
		FirstName:     in.FirstName,
		LastName:      in.LastName,
		Gender:        in.Gender,
		MaritalStatus: in.MaritalStatus,
		Email:         in.Email,
		PhoneNumber:   in.PhoneNumber,
	}
	if err := p.playerService.UpdatePlayer(in.Id, per); err != nil {
		return nil, fmt.Errorf("error updating player data %v", err)
	}
	//update field info
	field := models.FieldInfo{
		YearOfExperience:    in.YearOfExperience,
		NumberOfGoalsScored: in.NumberOfGoalsScored,
		JerseyNumber:        in.JerseyNumber,
		YearJoined:          in.YearJoined,
		PositionOnTheField:  in.PositionOnTheField,
		PlayerClubStatus:    in.PlayerStatus,
	}
	if err := p.playerService.UpdatePlayerWithFieldsInfo(in.Id, &field); err != nil {
		return nil, fmt.Errorf("error updating player field data %v", err)
	}

	//update address
	addr := &models.Address{
		Name:    in.Name,
		City:    in.City,
		ZipCode: in.ZipCode,
	}
	if err := p.playerService.UpdateAddress(in.Id, addr); err != nil {
		return nil, fmt.Errorf("error updating address data %v", err)
	}
	return &player.Empty{}, nil
}

func (p PlayerHandler) GetPlayerByJerseyNumber(ctx context.Context, in *player.GetPlayerUsingJerseyNumberRequest) (*player.GetPlayerByIdResponse, error) {
	//get player field info
	field, err := p.playerService.GetPlayerByJerseyNumber(in.JerseyNumber)
	if err != nil {
		return nil, fmt.Errorf("error getting player field data %v", err)
	}

	//get player personal info
	pers, err := p.playerService.GetPlayerById(field.PersonalInfoId)
	if err != nil {
		return nil, fmt.Errorf("error getting player personal data %v", err)
	}

	//get player address info
	add, err := p.playerService.GetAddressById(field.PersonalInfoId)
	if err != nil {
		return nil, fmt.Errorf("error getting player address info %v", err)
	}

	return &player.GetPlayerByIdResponse{
		FirstName:           pers.FirstName,
		LastName:            pers.LastName,
		Gender:              pers.Gender,
		MaritalStatus:       pers.MaritalStatus,
		Email:               pers.Email,
		PhoneNumber:         pers.PhoneNumber,
		YearOfExperience:    field.YearOfExperience,
		YearJoined:          field.YearJoined,
		PositionOnTheField:  field.PositionOnTheField,
		NumberOfGoalsScored: field.NumberOfGoalsScored,
		PlayerStatus:        field.PlayerClubStatus,
		JerseyNumber:        field.JerseyNumber,
		Name:                add.Name,
		City:                add.City,
		ZipCode:             add.ZipCode,
	}, nil

}

func (p PlayerHandler) DeletePlayer(ctx context.Context, in *player.DeletePlayerRequest) (*player.Empty, error) {
	playerId := in.Id
	if err := p.playerService.DeletePlayer(playerId); err != nil {
		return nil, fmt.Errorf("error deleting a player %v", err)
	}

	if err := p.playerService.DeletePlayerFieldInfo(playerId); err != nil {
		return nil, fmt.Errorf("error deleting a player field record %v", err)
	}

	if err := p.playerService.DeletePlayerAddress(playerId); err != nil {
		return nil, fmt.Errorf("error deleting a player address record %v", err)
	}
	return &player.Empty{}, nil
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

func generatePlayerId(length int) string {

	const (
		letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	)
	b := make([]byte, length)
	for i := range b {
		randomIndex := rand.Intn(len(letterBytes))
		b[i] = letterBytes[randomIndex]
	}

	return string(b)
}

/*query{
	paginatedOrganization(input:{filter:"",Page:10,Limit:1}){
   count
    organization{
      name
      email
      registration_number
      social_profile
    }
  }
  }

  mutation{
	login(input:{email:"amthetechguy@gmail.com",password:"harvest600"}){
	  access_token
	}
  }

  headers: {
	Authorization: 'Bearer ' + token
  }
*/
