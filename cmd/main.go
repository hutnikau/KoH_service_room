package main

import (
	"encoding/json"
	"fmt"

	"service-room/pkg/handlers"
	"service-room/pkg/infrastructure"
	"service-room/pkg/model"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rs/zerolog"
)

func main() {
	lambda.Start(handler)
}

type ErrorBody struct {
	ErrorMsg string `json:"error,omitempty"`
}

func handler(req events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	roomRepo := infrastructure.DynamodbRoomRepository{}
	h := handlers.NewHandler(roomRepo)

	fmt.Printf("%+v\n", req)

	switch req.RequestContext.RouteKey {
	case "GET /room":
		room := model.Room{}
		_ = json.Unmarshal([]byte(req.Body), &room)

		if id, ok := req.QueryStringParameters["id"]; ok {
			return h.GetUserRooms(id)
		}
	case "GET /user/{id}/rooms":
		if userId, ok := req.PathParameters["userId"]; ok {
			return h.GetRoom(userId)
		}
	case "POST /room":
		room := model.Room{}
		_ = json.Unmarshal([]byte(req.Body), &room)
		return h.CreateRoom(room)
	}

	return h.UnhandledMethod()
}
