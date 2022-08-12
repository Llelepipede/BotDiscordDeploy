package other

import (
	"golang-discord-bot/dataStruct"
	"reflect"
)

func FormDisplay(in []dataStruct.Api) string {
	Value := reflect.ValueOf(&in[0]).Elem()
	ret := ""
	for i := 0; i < Value.NumField(); i++ {
		ret += Value.Type().Field(i).Type.Name() + " | "
	}
	return ret
}
