package model

import (
	"encoding/json"

	"github.com/gofrs/uuid"
)

// BYTE converts object to byte
func BYTE(model interface{}) []byte {
	body, _ := json.Marshal(model)
	return body
}

// JSON converts object to json
func JSON(model interface{}) string {
	body, _ := json.Marshal(model)
	return string(body)
}

// Database Constant
const (
	MYSQLTimestampFormat = `2006-01-02 15:04:05`
)

// GenerateUUID ....
func GenerateUUID() string {
	uuid := uuid.Must(uuid.NewV4())
	return uuid.String()
}
