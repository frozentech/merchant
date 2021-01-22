package model

import "encoding/json"

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
