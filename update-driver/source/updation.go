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
	var info RequestBody
	err := json.Unmarshal([]byte(req.Body), &info)
	if err != nil {
		return ServerError(err)
	}
	app.UpdateDriver(info, req.QueryStringParameters["driverid"])
	mapd["error"] = false
	mapd["message"] = "Driver updated successfully"
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

func (app *App) UpdateDriver(data RequestBody, driverid string) {
	if data.FirstName != "" {
		app.DB.Debug().Exec(`UPDATE access_control.driver_details SET first_name=? where _id=?`, data.FirstName, driverid)
	}
	if data.LastName != "" {
		app.DB.Debug().Exec(`UPDATE access_control.driver_details SET last_name=? where _id=?`, data.LastName, driverid)
	}
	if data.Mobile != "" {
		app.DB.Debug().Exec(`UPDATE access_control.driver_details SET phone_no=? where _id=?`, data.Mobile, driverid)
	}
	if data.Password != "" {
		app.DB.Debug().Exec(`UPDATE access_control.driver_details SET password=? where _id=?`, data.Password, driverid)
	}
}
