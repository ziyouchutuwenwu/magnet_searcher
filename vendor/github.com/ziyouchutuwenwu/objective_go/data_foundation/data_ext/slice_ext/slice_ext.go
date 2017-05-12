package slice_ext

import (
	"reflect"
)

//start位置的会被删除，end位置的不会被删除
func SliceRemoveWithRange(slice interface{}, start, end int) interface{} {
	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		panic("wrong type")
	}

	cap := reflect.ValueOf(slice).Cap()

	startSlice := reflect.ValueOf(slice).Slice(0,start)
	endSlice := reflect.ValueOf(slice).Slice(end,cap)

	newSlice := reflect.AppendSlice(startSlice, endSlice)
	return newSlice.Interface()
}

func SliceRemoveAtIndex(slice interface{}, index int) interface{} {
	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		panic("wrong type")
	}

	newSlice := reflect.ValueOf(slice)
	if index < 0 || index >= newSlice.Len() {
		panic("out of bounds")
	}
	prev := newSlice.Index(index)
	for i := index + 1; i < newSlice.Len(); i++ {
		next := newSlice.Index(i)
		prev.Set(next)
		prev = next
	}
	return newSlice.Slice(0, newSlice.Len()-1).Interface()
}

func SliceReplaceAtIndex(slice interface{}, index int, object interface{}) interface{} {
	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		panic("wrong type")
	}
	newSlice := reflect.ValueOf(slice)
	if index < 0 || index >= newSlice.Len() {
		panic("out of bounds")
	}
	element := newSlice.Index(index)
	element.Set(reflect.ValueOf(object))
	return newSlice.Interface()
}

func SliceAppendObject(slice interface{}, object interface{}) interface{} {
	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		panic("wrong type")
	}
	newSlice := reflect.Append(reflect.ValueOf(slice), reflect.ValueOf(object))
	return newSlice.Interface()
}

func SliceAppendSlice(slice interface{}, sliceToAppend interface{}) interface{} {
	if reflect.TypeOf(slice).Kind() != reflect.Slice || reflect.TypeOf(sliceToAppend).Kind() != reflect.Slice {
		panic("wrong type")
	}
	newSlice := reflect.ValueOf(slice)
	newSlice = reflect.AppendSlice(newSlice, reflect.ValueOf(sliceToAppend))
	return newSlice.Interface()
}

func IsSliceEqual(slice1 interface{}, slice2 interface{}) bool {
	if reflect.TypeOf(slice1).Kind() != reflect.Slice || reflect.TypeOf(slice2).Kind() != reflect.Slice {
		panic("wrong type")
	}
	return reflect.DeepEqual(slice1, slice2)
}
