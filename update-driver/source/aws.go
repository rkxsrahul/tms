package source

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// InitSession create a session for login credentials
func Init() (*session.Session, error) {
	var sess *session.Session
	var err error
	sess, err = session.NewSession(&aws.Config{
		Region: aws.String(REGION_NAME),
	})
	return sess, err
}

// SSMvalue is used to get value using name
func SSMvalue(name string) string {
	sess, err := Init()
	if err != nil {
		log.Println(err)
		return ""
	}
	//connect to ssm service
	ssmsvc := ssm.New(sess, aws.NewConfig().WithRegion(REGION_NAME))
	param, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(name),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		log.Println(err)
		return ""
	}

	return *param.Parameter.Value
}

func (app *App) AddUser(email, name, password, role string) (error, string) {
	authTry := &cognito.AdminCreateUserInput{
		TemporaryPassword: aws.String(password),
		UserAttributes: []*cognito.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(email),
			}, {
				Name:  aws.String("name"),
				Value: aws.String(name),
			}, {
				Name:  aws.String("email_verified"),
				Value: aws.String("true"),
			},
		},
		UserPoolId: aws.String(SSMvalue(COGNITO_USER_POOL_ID)),
		Username:   aws.String(email),
	}

	_, err := app.CognitoClient.AdminCreateUser(authTry)
	if err != nil {
		log.Println(err)
		return err, ""
	}
	userdetails := &cognito.AdminGetUserInput{
		UserPoolId: aws.String(SSMvalue(COGNITO_USER_POOL_ID)),
		Username:   aws.String(email),
	}
	userdata, err := app.CognitoClient.AdminGetUser(userdetails)
	if err != nil {
		log.Println(err)
		return err, ""
	}
	return nil, *userdata.Username
}

func computeSecretHash(clientSecret string, username string, clientId string) string {
	mac := hmac.New(sha256.New, []byte(clientSecret))
	mac.Write([]byte(username + clientId))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
