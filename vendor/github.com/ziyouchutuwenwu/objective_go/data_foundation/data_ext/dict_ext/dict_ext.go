package dict_ext

import (
	"reflect"
)

func IsDictionaryEqual(dict1 interface{}, dict2 interface{}) bool {
	if reflect.TypeOf(dict1).Kind() != reflect.Map || reflect.TypeOf(dict2).Kind() != reflect.Map {
		panic("wrong type")
	}
	return reflect.DeepEqual(dict1, dict2)
}
