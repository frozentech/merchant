package model

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

// Model Constant
const (
	DBName               = "pace_merchant"
	MYSQLTimestampFormat = `2006-01-02 15:04:05`
)

// DB ...
var (
	DB        *sqlx.DB
	SQSClient *sqs.SQS
	IsClose   bool
	Retry     = 50
)

// Contains ....
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// ConnectDB ...
func ConnectDB() (*sqlx.DB, error) {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("ENV_DB_USERNAME"),
		os.Getenv("ENV_DB_PASSWORD"),
		os.Getenv("ENV_DB_HOST"),
		os.Getenv("ENV_DB_PORT"),
		os.Getenv("ENV_DB_NAME"))

	dbConn, err := sqlx.Connect(os.Getenv("ENV_DB_ENGINE"), connection)
	if err != nil {
		fmt.Println(connection)
	}

	return dbConn, err
}

// JSON ...
func JSON(model interface{}) string {
	body, _ := json.Marshal(model)
	return string(body)
}

// BYTE ...
func BYTE(model interface{}) []byte {
	body, _ := json.Marshal(model)
	return body
}

// GenerateUUID ....
func GenerateUUID() string {
	uuid := uuid.Must(uuid.NewV4())
	return uuid.String()
}
