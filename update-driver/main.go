package main

import (
	"errors"
	"log"
	"strings"

	"bitbucket.org/tms/typ23-all-apis/user-creation/source"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Route(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println(req.Body)
	db, err := source.PostgresDB()
	if err != nil {
		return source.ServerError(err)
	}
	app := source.App{
		DB: db,
	}
	status, _ := app.CheckJWTToken(strings.Split(req.Headers["Authorization"], " ")[1])
	if status != "Authorized" {
		return source.InternalError(errors.New("user is not authorized"))
	}
	return app.CreateUser(req)
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	//Invoke lambda
	lambda.Start(Route)
}
