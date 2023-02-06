package source

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// UpdateUser function will update the user firstname and rolename
func (app *App) CreateUser(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	mapd := make(map[string]interface{})
	driverid := req.QueryStringParameters["driverId"]
	err := app.DeleteDriver(driverid)
	if err != nil {
		return ServerError(err)
	}
	mapd["error"] = false
	mapd["message"] = "driver is deleted"
	bytedata, err := json.Marshal(mapd)
	if err != nil {
		log.Print(err)
		return ServerError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(bytedata),
	}, nil
}

func (app *App) DeleteDriver(driverid string) error {
	err := app.DB.Debug().Exec(`DELETE FROM access_control.driver_details WHERE _id=?`, driverid).Error
	if err != nil {
		return err
	}
	return nil
}
