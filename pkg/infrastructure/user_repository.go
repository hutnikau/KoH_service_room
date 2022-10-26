package infrastructure

import (
	"encoding/json"
	"fmt"
	"os"
	"service-room/pkg/handlers"
	"service-room/pkg/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type LambdaUserRepository struct {
}

type getItemsRequest struct {
	Message string `json:"message"`
}

func NewUserRepository() handlers.UserRepository {
	r := LambdaUserRepository{}
	return r
}

func (l LambdaUserRepository) GetUserById(id string) (model.User, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	client := lambda.New(sess, &aws.Config{Region: aws.String("eu-central-1")})
	request := getItemsRequest{
		Message: "test",
	}

	payload, _ := json.Marshal(request)

	result, err := client.Invoke(&lambda.InvokeInput{FunctionName: aws.String("KoH_dev_auth"), Payload: payload})
	if err != nil {
		fmt.Printf("%+v\n", err)
		fmt.Println("Error calling KoH_prod_auth")
		os.Exit(0)
	}

	fmt.Printf("%+v\n", result)

	return model.User{}, nil
}
