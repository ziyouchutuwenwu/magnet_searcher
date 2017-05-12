package array

import (
	"github.com/mitchellh/copystructure"
	"github.com/ziyouchutuwenwu/objective_go/data_foundation/data_ext/slice_ext"
)

type Array []interface{}

func Create() Array {
	array := new(Array)
	return *array
}

func (this *Array) Init() {
}

func (self Array) Len() int {
	return len(self)
}

func (this *Array) AddObject(object interface{}) Array {
	*this = slice_ext.SliceAppendObject(*this, object).(Array)
	return *this
}

func (this *Array) AddObjects(objects Array) Array {
	*this = slice_ext.SliceAppendSlice(*this, objects).(Array)
	return *this
}

func (this *Array) RemoveObjectAtIndex(index int) Array {
	*this = slice_ext.SliceRemoveAtIndex(*this, index).(Array)
	return *this
}

func (this *Array) RemoveObject(object interface{}) Array {
	index := this.FindObject(object)
	if index >= 0 {
		this.RemoveObjectAtIndex(index)
	}
	return *this
}

func (this *Array) RemoveObjectWithRange(start int, end int) Array {
	*this = slice_ext.SliceRemoveWithRange(*this, start, end).(Array)
	return *this
}

func (this *Array) RemoveAllObjects() Array {
	*this = slice_ext.SliceRemoveWithRange(*this, 0, this.Len()).(Array)
	return *this
}

func (this *Array) ReplaceObjectAtIndex(index int, object interface{}) Array {
	*this = slice_ext.SliceReplaceAtIndex(*this, index, object).(Array)
	return *this
}

func (this *Array) Clone() Array {
	newArray := Create()
	newArray.Init()

	for i:= 0;i < this.Len();i++{
		object := this.GetObjectAtIndex(i)

		if nil != object{
			newObject,_ := copystructure.Copy(object)
			newArray.AddObject(newObject)
		}
	}

	return newArray
}

func (self Array) FindObject(object interface{}) int {
	index := -1

	for i := 0; i < self.Len(); i++ {
		if object == self.GetObjectAtIndex(i) {
			index = i
			break
		}
	}
	return index
}

func (self Array) GetObjectAtIndex(index int) interface{} {
	return self[index]
}

func (self Array) LastObject() interface{} {
	return self.GetObjectAtIndex(self.Len() - 1)
}

func (self Array) IsEqualToArray(arrayToCompare Array) bool {
	return slice_ext.IsSliceEqual(self, arrayToCompare)
}