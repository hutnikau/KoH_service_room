package handlers

import (
	"net/http"
	"service-room/pkg/model"

	"github.com/aws/aws-lambda-go/events"
)

type Handler struct {
	roomRepo RoomRepository
}

type ErrorBody struct {
	ErrorMsg string `json:"error,omitempty"`
}

func NewHandler(roomRepo RoomRepository) Handler {
	return Handler{
		roomRepo: roomRepo,
	}
}

func (h Handler) CreateRoom(room model.Room) (*events.APIGatewayV2HTTPResponse, error) {
	newRoom, err := h.roomRepo.CreateRoom(room)
	if err != nil {
		return ApiResponse(http.StatusBadRequest, ErrorBody{err.Error()}), err
	}
	return ApiResponse(http.StatusOK, newRoom), nil
}

func (h Handler) GetRoom(id string) (*events.APIGatewayV2HTTPResponse, error) {
	room, err := h.roomRepo.GetRoomById(id)
	if err != nil {
		return ApiResponse(http.StatusNotFound, ErrorBody{"Room not found"}), err
	}
	return ApiResponse(http.StatusOK, room), nil
}

func (h Handler) GetUserRooms(userId string) (*events.APIGatewayV2HTTPResponse, error) {
	rooms, err := h.roomRepo.GetRoomsByUserId(userId)
	if err != nil {
		return ApiResponse(http.StatusBadRequest, ErrorBody{err.Error()}), err
	}
	return ApiResponse(http.StatusOK, rooms), nil
}

func (h Handler) UnhandledMethod() (*events.APIGatewayV2HTTPResponse, error) {
	return ApiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed), nil
}
