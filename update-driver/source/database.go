package source

import (
	"encoding/json"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func UnMarshalDbCredentials() DatabaseCredentialsStore {
	var info DatabaseCredentialsStore
	//get value from name
	ssmvalue := SSMvalue(TMS_DB)
	//get data from json string
	err := json.Unmarshal([]byte(ssmvalue), &info)
	if err != nil {
		log.Println(err)
	}
	return info
}

// PostgresDB -> postgres database connection
func PostgresDB() (*gorm.DB, error) {
	dbcredentials := UnMarshalDbCredentials()
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		//database host
		dbcredentials.Host,
		//database port
		dbcredentials.Port,
		//database user
		dbcredentials.User,
		//database pass
		dbcredentials.Password,
		//database name
		dbcredentials.Database,
		//ssl
		"disable")), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return db, err
	}
	return db, nil
}
