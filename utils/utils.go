package utils

import "encoding/json"

func Encode(data interface{}) string {
	dataByte, _ := json.Marshal(data)
	return string(dataByte)
}
