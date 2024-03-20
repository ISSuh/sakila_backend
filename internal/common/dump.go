package common

import "encoding/json"

func Dump(data any) string {
	str, _ := json.Marshal(data)
	return string(str)
}
