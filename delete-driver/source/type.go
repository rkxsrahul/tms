package source

import (
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"gorm.io/gorm"
)

type App struct {
	// Database client
	DB            *gorm.DB
	CognitoClient *cognito.CognitoIdentityProvider
}

type DatabaseCredentialsStore struct {
	Database string `json:"database"`
	Host     string `json:"host"`
	Password string `json:"password"`
	Port     int    `json:"port"`
	User     string `json:"user"`
}

type Users struct {
	Username      string `json:"username"`
	Email         string `json:"email"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	PhoneNo       string `json:"phone_no"`
	RoleID        int    `json:"role_id"`
	IsActive      bool   `json:"is_active"`
	CreatedBy     int    `json:"created_by"`
	UpdatedBy     int    `json:"updated_by"`
	UserCognitoID string `json:"user_cognito_id"`
}
