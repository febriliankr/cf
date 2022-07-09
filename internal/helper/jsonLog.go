package helper

import "encoding/json"

func JSONLog(data interface{}) string {
	j, _ := json.Marshal(data)

	return string(j)
}
