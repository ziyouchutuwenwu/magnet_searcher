package slice_ext

import(
	"reflect"
)

func DeDuplicate(a interface{}) (ret []interface{}) {
	value := reflect.ValueOf(a)
	for i := 0; i < value.Len(); i++ {
		if i > 0 && reflect.DeepEqual(value.Index(i-1).Interface(), value.Index(i).Interface()) {
			continue
		}
		ret = append(ret, value.Index(i).Interface())
	}
	return ret
}