package main

import (
	"fmt"
	"lambda-func/app"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	UserName string `json:"username"`
}

func HandleRequest(event MyEvent) (string, error) {
	if event.UserName == "" {
		return "", fmt.Errorf("username cannot be empty")
	}
	return fmt.Sprintf("Succescully called by - %s", event.UserName), nil
}

func main() {
	myApp := app.NewApp()
	lambda.Start(func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		switch request.Path {
		case "/register":
			return myApp.ApiHandler.RegisterUserHandler(request)
		case "/login":
			return myApp.ApiHandler.LoginUser(request)
		default:
			return events.APIGatewayProxyResponse{
				Body:       "not found",
				StatusCode: http.StatusNotFound,
			}, nil
		}
	})
}
