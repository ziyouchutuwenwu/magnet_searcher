package dictionary

import (
	"github.com/mitchellh/copystructure"
	"github.com/ziyouchutuwenwu/objective_go/data_foundation/data_ext/dict_ext"
	"reflect"
)

type Dictionary map[interface{}]interface{}

func Create() Dictionary {
	dict := new(Dictionary)
	return *dict
}

func (this *Dictionary) Init() {
	if reflect.TypeOf(*this).Kind() != reflect.Map {
		panic("wrong type")
	}
	*this = reflect.MakeMap(reflect.TypeOf(*this)).Interface().(Dictionary)
}

func (self Dictionary) Len() int {
	return len(self)
}

func (this *Dictionary) GetObjectForKey(key interface{}) interface{} {
	return (*this)[key]
}

func (this *Dictionary) SetObjectForKey(object interface{}, key interface{}) {
	(*this)[key] = object
}

func (this *Dictionary) RemoveObjectForKey(key interface{}) {
	delete(*this, key)
}

func (self Dictionary) Clone() Dictionary {
	newDict := Create()
	newDict.Init()

	for key, value := range self {
		if nil == key || nil == value{
			continue
		}
		newKey, _ := copystructure.Copy(key)
		newValue, _ := copystructure.Copy(value)
		newDict.SetObjectForKey(newValue,newKey)
	}

	return newDict
}

func (self Dictionary) IsEqualToDictionary(dictToCompare Dictionary) bool {
	return dict_ext.IsDictionaryEqual(self, dictToCompare)
}
