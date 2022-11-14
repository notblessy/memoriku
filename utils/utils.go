package utils

import "encoding/json"

func Encode(data interface{}) string {
	dataByte, _ := json.Marshal(data)
	return string(dataByte)
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
