package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"service-room/pkg/model"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockedRoomRepository struct {
	mock.Mock
}

// type RoomRepository interface {
// 	GetRoomById(id string) (model.Room, error)
// 	CreateRoom(room model.Room) (model.Room, error)
// 	GetRoomsByUserId(userId string) ([]model.Room, error)
// }

func (m MockedRoomRepository) GetRoomById(id string) (model.Room, error) {
	args := m.Called(id)
	return args.Get(0).(model.Room), args.Error(1)
}
func (m MockedRoomRepository) CreateRoom(room model.Room) (model.Room, error) {
	args := m.Called(room)
	return args.Get(0).(model.Room), args.Error(1)
}
func (m MockedRoomRepository) GetRoomsByUserId(userId string) ([]model.Room, error) {
	args := m.Called(userId)
	return args.Get(0).([]model.Room), args.Error(1)
}

func TestGetRoomByIdNotFound(test *testing.T) {
	roomRepo := MockedRoomRepository{}
	roomRepo.On("GetRoomById", "42").Return(model.Room{}, errors.New("could not find room"))

	handler := NewHandler(roomRepo)

	response, _ := handler.GetRoom("42")

	if response.StatusCode != 404 {
		test.Errorf("Expected status code 404, %d given", response.StatusCode)
	}

	expectedBody := `{"error":"Room not found"}`
	if response.Body != expectedBody {
		test.Errorf("Wrong body given: '%s'", response.Body)
	}
}

func TestGetRoomById(test *testing.T) {
	roomRepo := MockedRoomRepository{}
	roomRepo.On("GetRoomById", "42").Return(model.Room{
		Id:   "42",
		Name: "Red room",
		Owner: model.User{
			Id:    "user_42",
			Login: "user_42_login",
		},
		Users: []model.User{
			model.User{
				Id:    "user_42",
				Login: "user_42_login",
			},
			model.User{
				Id:    "user_146",
				Login: "user_146_login",
			},
		},
	}, nil)

	handler := NewHandler(roomRepo)

	response, _ := handler.GetRoom("42")
	expectedBody := `{"id":"42","login":"Red room","owner":{"id":"user_42","login":"user_42_login"},"users":[{"id":"user_42","login":"user_42_login"},{"id":"user_146","login":"user_146_login"}]}`

	if response.StatusCode != 200 {
		test.Errorf("Expected status code 200, %d given", response.StatusCode)
	}
	if response.Body != expectedBody {
		test.Errorf("Expected body worng: '%s'", response.Body)
	}
}

func TestCreateRoomError(test *testing.T) {
	roomRepo := MockedRoomRepository{}
	newRoom := model.Room{
		Name: "42",
	}
	roomRepo.On("CreateRoom", newRoom).Return(model.Room{}, errors.New("bad error message"))

	handler := NewHandler(roomRepo)

	response, _ := handler.CreateRoom(newRoom)

	if response.StatusCode != http.StatusBadRequest {
		test.Errorf("Expected status code 400, %d given", response.StatusCode)
	}
	fmt.Printf("%+v\n", newRoom)

	expectedBody := `{"error":"bad error message"}`
	if response.Body != expectedBody {
		test.Errorf("Wrong body given: '%s'", response.Body)
	}
}
