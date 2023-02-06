package source

import (
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
)

const (
	serviceName       = "Branch Configuration"
	servicePermission = 8
)

type Permission struct {
	ReadAccess   int `json:"read_access"`
	UpdateAccess int `json:"update_access"`
	WriteAccess  int `json:"write_access"`
	DeleteAccess int `json:"delete_access"`
}

// CheckJWTToken is used to validate the JWT token
func (app *App) CheckJWTToken(token string) (string, string) {
	jwttoken, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	claims, _ := jwttoken.Claims.(jwt.MapClaims)
	email := fmt.Sprintf("%v", claims["email"])
	cognitoid := fmt.Sprintf("%v", claims["cognito:username"])
	log.Println(app.CheckUserPermission(email, serviceName).WriteAccess)
	if app.CheckUserPermission(email, serviceName).DeleteAccess == servicePermission {
		return "Authorized", cognitoid
	}
	return "Not Authorized", ""
}

func (app *App) CheckUserPermission(email, servicename string) Permission {
	var permission Permission
	app.DB.Debug().Raw(`SELECT * FROM access_control.role_service_mapping 
	 where role_id=(SELECT role_id FROM access_control.users where email=?) 
	 and service_id=(SELECT service_id  FROM access_control.services where service_name=?)`, email, servicename).Scan(&permission)
	return permission
}
