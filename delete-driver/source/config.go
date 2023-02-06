package source

import "os"

var (
	TMS_DB               = os.Getenv("TMS_DB")
	REGION_NAME          = os.Getenv("REGION_NAME")
	COGNITO_USER_POOL_ID = os.Getenv("COGNITO_USER_POOL_ID")
)
