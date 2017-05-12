package array

import (
	"testing"
)

type A struct {
	Name string
}

func TestArray(t *testing.T) {
	oldArray := Create()
	oldArray.Init()

	oldArray.AddObject(10)
	oldArray.AddObject(20)
	oldArray.AddObject(30)
	oldArray.AddObject("223")

	object := new(A)
	oldArray.AddObject(object)
	t.Log("oldArray", oldArray)

	newArray := oldArray.Clone()
	t.Log("cloneArray", newArray)

	newArray.RemoveObjectAtIndex(1)
	t.Log("removedArray", newArray)

	t.Log("isEqual?", oldArray.IsEqualToArray(newArray))

	index := oldArray.FindObject(object)
	t.Log("find index?", index)

	t.Log("oldArray", oldArray)
	oldArray.ReplaceObjectAtIndex(0, object)
	t.Log("replaced at 0", oldArray)

	arrayToAppend := Create()
	arrayToAppend.AddObject(object)
	arrayToAppend.AddObject(object)
	arrayToAppend.AddObject(object)

	oldArray.AddObjects(arrayToAppend)
	t.Log("appended array", oldArray)

	oldArray.RemoveObjectWithRange(2, 5)
	t.Log("range removed array", oldArray)
}
