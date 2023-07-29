package server

import (
	"context"
	"testing"

	"github.com/DecodeWorms/server-contract/models"
	"github.com/DecodeWorms/sv.player/mocks"
	"github.com/DecodeWorms/sv.player/pb/protos/pb/player"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreatePlayer(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrl.Finish()

	user := &models.PersonalInfo{
		FirstName:     "Bola",
		LastName:      "Yinka",
		Gender:        "female",
		MaritalStatus: "single",
		Email:         "test@gmail.com",
		PhoneNumber:   "09050095721",
	}

	storeMock := mocks.NewMockPlayerStore(ctrl)
	storeMock.EXPECT().GetPlayerByEmail(gomock.Any()).Return(user, nil).Times(1)
	storeMock.EXPECT().GetPlayerByPhoneNumber("09050095721").Return(user, nil).Times(1)
	storeMock.EXPECT().CreatePlayer(gomock.Any()).Return(nil).Times(1)
	handler, _ := NewPlayerHandler(storeMock)
	_, err := handler.CreatePlayer(context.Background(), &player.CreatePlayerRequest{
		Id:            "test-id",
		FirstName:     "John",
		LastName:      "Doe",
		Gender:        "male",
		PhoneNumber:   "090000000000",
		Email:         "test@mail.com",
		MaritalStatus: "single",
	})
	assert.EqualError(t, err, "error email already exist <nil>")
}

func TestDeletePlayer(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrl.Finish()

	id := "test-123"
	storemock := mocks.NewMockPlayerStore(ctrl)
	storemock.EXPECT().DeletePlayer(id).Return(nil).Times(1)

	handler, _ := NewPlayerHandler(storemock)
	err := handler.playerService.DeletePlayer(id)
	assert.Nil(t, err)
}

func TestGetplayerById(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrl.Finish()

	resp := &models.PersonalInfo{
		FirstName:     "John",
		LastName:      "Doe",
		Gender:        "male",
		MaritalStatus: "married",
		Email:         "johndoe@mail.com",
		PhoneNumber:   "090000000000",
	}

	respo := &models.FieldInfo{
		YearOfExperience:    "22",
		NumberOfGoalsScored: "10",
		JerseyNumber:        "10",
		YearJoined:          "2018-3-22",
		PersonalInfoId:      "Striker",
		PlayerClubStatus:    "available",
	}

	storemock := mocks.NewMockPlayerStore(ctrl)
	storemock.EXPECT().GetPlayerById(gomock.Any()).Return(resp, nil).Times(1)
	storemock.EXPECT().GetPlayerWithFieldsInfoById(gomock.Any()).Return(respo, nil).Times(1)
	handler, _ := NewPlayerHandler(storemock)
	res, err := handler.playerService.GetPlayerById("id")
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestGetPlayerByPhoneNumber(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrl.Finish()

	resp := &models.PersonalInfo{
		FirstName:     "John",
		LastName:      "Doe",
		Gender:        "male",
		MaritalStatus: "married",
		Email:         "johndoe@mail.com",
		PhoneNumber:   "090000000000",
	}

	respo := &models.FieldInfo{
		YearOfExperience:    "22",
		NumberOfGoalsScored: "10",
		JerseyNumber:        "10",
		YearJoined:          "2018-3-22",
		PersonalInfoId:      "Striker",
		PlayerClubStatus:    "available",
	}

	storemock := mocks.NewMockPlayerStore(ctrl)
	storemock.EXPECT().GetPlayerByPhoneNumber(gomock.Any()).Return(resp, nil).Times(1)
	storemock.EXPECT().GetPlayerWithFieldsInfoById(gomock.Any()).Return(respo, nil).Times(1)
	handler, _ := NewPlayerHandler(storemock)
	rest, err := handler.playerService.GetPlayerByPhoneNumber("test-123")
	assert.Nil(t, err)
	assert.NotNil(t, rest)

}

/*func TestUpdatePlayerWithFieldsInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	storeMock := mocks.NewMockPlayerStore(ctrl)
	mockRequest := &player.UpdatePlayerWithFieldDataRequest{
		PersonalInfoId:      "personal-test-id-123",
		YearOfExperience:    "22 years",
		NumberOfGoals:       33,
		JerseyNumber:        10,
		YearJoined:          "2019-10-10",
		PositionOnTheField:  "Striker",
		NumberOfGoalsScored: 33,
		PlayerStatus:        "Not available",
	}


	storeMock.EXPECT().UpdatePlayerWithFieldsInfo(mockRequest.PersonalInfoId, gomock.Any()).Return(nil).Times(1)
	handler, _ := New(storeMock)
	_, err := handler.UpdatePayerWithFieldData(context.Background(), &player.UpdatePlayerWithFieldDataRequest{
		PersonalInfoId:      "personal-test-id-123",
		YearOfExperience:    "22 years",
		NumberOfGoals:       33,
		JerseyNumber:        10,
		YearJoined:          "2019-10-10",
		PositionOnTheField:  "Striker",
		NumberOfGoalsScored: 33,
		PlayerStatus:        "Not available",
	})
	assert.Nil(t, err)
}
*/
