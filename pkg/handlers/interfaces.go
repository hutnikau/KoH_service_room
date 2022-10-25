package handlers

import "service-room/pkg/model"

type RoomRepository interface {
	GetRoomById(id string) (model.Room, error)
	CreateRoom(room model.Room) (model.Room, error)
	GetRoomsByUserId(userId string) ([]model.Room, error)
}
