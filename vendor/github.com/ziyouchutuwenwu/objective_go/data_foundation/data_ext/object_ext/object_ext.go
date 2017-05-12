package object_ext

import (
	"reflect"
)

func IsObjectEqual(object1 interface{}, object2 interface{}) bool {
	return reflect.DeepEqual(object1, object2)
}
