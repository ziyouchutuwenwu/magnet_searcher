package object

import (
	"github.com/mitchellh/copystructure"
	"github.com/ziyouchutuwenwu/objective_go/data_foundation/data_ext/object_ext"
	"reflect"
)

type Object struct {
	Value interface{}
}

func Create() Object {
	object := new(Object)
	return *object
}

func (this *Object) Init() {
}

func (self Object) Clone() Object {
	newObject := Create()

	if nil != self.Value{
		newObject.Value, _ = copystructure.Copy(self.Value)
	}
	return newObject
}

func (self Object) GetValueType() reflect.Type {
	return reflect.TypeOf(self.Value)
}

func (self Object) GetValueTypeName() string {
	return reflect.TypeOf(self.Value).Name()
}

func (self Object) IsEqualToObject(objectToCompare Object) bool {
	return object_ext.IsObjectEqual(self, objectToCompare)
}
