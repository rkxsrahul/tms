package source

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

// Add a helper for handling errors. This logs any error to os.Stderr
// and returns a 500 Internal Server Error response that the AWS API
// Gateway understands.
func ServerError(err error) (events.APIGatewayProxyResponse, error) {
	mapd := make(map[string]interface{})
	mapd["error"] = true
	mapd["message"] = err.Error()
	byteData, err := json.Marshal(mapd)
	if err != nil {
		return ServerError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 400,
		Body:       string(byteData),
	}, nil
}

func InternalError(err error) (events.APIGatewayProxyResponse, error) {
	mapd := make(map[string]interface{})
	mapd["error"] = true
	mapd["message"] = err.Error()
	byteData, err := json.Marshal(mapd)
	if err != nil {
		return ServerError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       string(byteData),
	}, nil
}
