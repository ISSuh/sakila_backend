package common

import (
	"encoding/json"
	"reflect"
)

func Dump(data any) string {
	str, _ := json.Marshal(data)
	return string(str)
}

func GetType(myvar any) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}
