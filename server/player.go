package server

import (
	"context"
	"fmt"
	"math/rand"

	//pr "github.com/DecodeWorms/messaging-protocol"
	"github.com/DecodeWorms/server-contract/models"
	db "github.com/DecodeWorms/sv.player/db/postgres"
	"github.com/DecodeWorms/sv.player/pb/protos/pb/player"
	"google.golang.org/grpc"
)

type PlayerHandler struct {
	playerService db.PostgresStore
	//store         pr.PulsarStore
	//add the logger
}

func New(p db.PostgresStore) (PlayerHandler, error) {
	return PlayerHandler{
		playerService: p,
	}, nil
}

var _ player.PlayerServiceClient = PlayerHandler{}

func (p PlayerHandler) CreatePlayerWithFieldData(ctx context.Context, in *player.CreatePlayerFieldRequest, op ...grpc.CallOption) (*player.Empty, error) {
	data := models.FieldInfo{
		PersonalInfoId:      in.PlayerId,
		YearOfExperience:    in.YearOfExperience,
		NumberOfGoalsScored: int(in.NumberOfGoals),
		JerseyNumber:        int(in.JerseyNumber),
		YearJoined:          in.YearJoined,
		PositionOnTheField:  in.PositionOnTheField,
		PlayerClubStatus:    in.PlayerStatus,
	}
	if err := p.playerService.CreatePlayerWithFieldsData(data); err != nil {
		return nil, fmt.Errorf("error creating player %v", err)
	}
	return &player.Empty{}, nil
}

// Complete me , am yet to be fully implemented
func (p PlayerHandler) GetPlayerByJerseyNumber(ctx context.Context, in *player.GetPlayerUsingJerseyNumberRequest, op ...grpc.CallOption) (*player.GetPlayerByIdResponse, error) {
	return nil, nil
}

func (p PlayerHandler) UpdatePlayerWithFieldsInfo(ctx context.Context, in *player.UpdatePlayerWithFieldDataRequest) (*player.Empty, error) {
	playerId := in.PersonalInfoId
	data := &models.FieldInfo{
		YearOfExperience:    in.YearOfExperience,
		NumberOfGoalsScored: int(in.NumberOfGoalsScored),
		JerseyNumber:        int(in.JerseyNumber),
		YearJoined:          in.YearJoined,
		PositionOnTheField:  in.PositionOnTheField,
		PlayerClubStatus:    in.PlayerStatus,
	}
	if err := p.playerService.UpdatePlayerWithFieldsInfo(playerId, data); err != nil {
		return nil, fmt.Errorf("error updating player record %v", err)
	}
	return &player.Empty{}, nil
}

// Delete me , because there is UpdatePlayerWithFieldsInfo() which also do update
func (p PlayerHandler) UpdatePayerWithFieldData(ctx context.Context, in *player.UpdatePlayerWithFieldDataRequest, op ...grpc.CallOption) (*player.Empty, error) {
	return nil, nil
}

func (p PlayerHandler) DeletePlayer(ctx context.Context, in *player.DeletePlayerRequest, op ...grpc.CallOption) (*player.Empty, error) {
	playerId := in.Id
	if err := p.playerService.DeletePlayer(playerId); err != nil {
		return nil, fmt.Errorf("error deleting a player %v", err)
	}
	return &player.Empty{}, nil
}
func (p PlayerHandler) CreatePlayer(ctx context.Context, in *player.CreatePlayerRequest, op ...grpc.CallOption) (*player.Empty, error) {
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

func (p PlayerHandler) GetPlayerById(ctx context.Context, in *player.GetPlayerByIdRequest, op ...grpc.CallOption) (*player.GetPlayerByIdResponse, error) {
	playerId := in.Id
	res, err := p.playerService.GetPlayerById(playerId)
	if err != nil {
		return nil, fmt.Errorf("error getting player info %v", err)
	}

	field, err := p.playerService.GetPlayerWithFieldsInfoById(playerId)
	if err != nil {
		return nil, fmt.Errorf("error getting player field info %v", err)
	}
	return &player.GetPlayerByIdResponse{
		FirstName:           res.FirstName,
		LastName:            res.LastName,
		Gender:              res.Gender,
		PhoneNumber:         res.PhoneNumber,
		MaritalStatus:       res.MaritalStatus,
		Email:               res.Email,
		YearOfExperience:    field.YearOfExperience,
		NumberOfGoalsScored: int32(field.NumberOfGoalsScored),
		JerseyNumber:        int32(field.JerseyNumber),
		YearJoined:          field.YearJoined,
		PositionOnTheField:  field.PositionOnTheField,
		PlayerStatus:        field.PlayerClubStatus,
	}, nil
}

func (p PlayerHandler) GetPlayerByPhoneNumber(ctx context.Context, in *player.GetPlayerByPhoneNumberRequest, op ...grpc.CallOption) (*player.GetPlayerByIdResponse, error) {
	playerId := in.PhoneNumber
	res, err := p.playerService.GetPlayerByPhoneNumber(playerId)
	if err != nil {
		return nil, fmt.Errorf("error getting player data %v", err)
	}
	field, err := p.playerService.GetPlayerWithFieldsInfoById(playerId)
	if err != nil {
		return nil, fmt.Errorf("error getting player field info %v", err)
	}
	return &player.GetPlayerByIdResponse{
		FirstName:           res.FirstName,
		LastName:            res.LastName,
		Gender:              res.Gender,
		PhoneNumber:         res.PhoneNumber,
		MaritalStatus:       res.MaritalStatus,
		Email:               res.Email,
		YearOfExperience:    field.YearOfExperience,
		NumberOfGoalsScored: int32(field.NumberOfGoalsScored),
		JerseyNumber:        int32(field.JerseyNumber),
		YearJoined:          field.YearJoined,
		PositionOnTheField:  field.PositionOnTheField,
		PlayerStatus:        field.PlayerClubStatus,
	}, nil
}

func (p PlayerHandler) UpdatePlayerPersonalInfo(ctx context.Context, in *player.UpdatePlayerPersonalInfoRequest, op ...grpc.CallOption) (*player.Empty, error) {
	data := &models.PersonalInfo{
		FirstName:     in.FirstName,
		LastName:      in.LastName,
		Gender:        in.Gender,
		MaritalStatus: in.MaritalStatus,
		Email:         in.Email,
		PhoneNumber:   in.PhoneNumber,
	}
	playerId := in.Id
	if err := p.playerService.UpdatePlayer(playerId, data); err != nil {
		return nil, fmt.Errorf("error updating player data %v", err)
	}
	return &player.Empty{}, nil
}

// Complete me , am yet to be fully implemented
func (p PlayerHandler) UpdateAddress(ctx context.Context, in *player.UpdateAddressRequest, op ...grpc.CallOption) (*player.Empty, error) {
	return nil, nil
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
