package infrastructure

import "service-room/pkg/model"

type DynamodbRoomRepository struct {
}

func (m DynamodbRoomRepository) GetRoomById(id string) (model.Room, error) {
	return model.Room{}, nil
}

func (m DynamodbRoomRepository) CreateRoom(room model.Room) (model.Room, error) {
	return room, nil
}

func (m DynamodbRoomRepository) GetRoomsByUserId(userId string) ([]model.Room, error) {
	return []model.Room{}, nil
}
