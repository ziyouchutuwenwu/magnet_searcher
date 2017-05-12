package dictionary

import "testing"

type A struct {
	Name string
}

func TestDictionary(t *testing.T) {
	var oldDict = Create()
	oldDict.Init()

	oldDict.SetObjectForKey("23123", 123)
	oldDict.SetObjectForKey("kick", "ggg")

	object := new(A)
	object.Name = "Name"

	oldDict.SetObjectForKey(object, "object")

	newDict := oldDict.Clone()
	t.Log(oldDict)

	newDict.RemoveObjectForKey("ggg")

	t.Log(oldDict, newDict)
}
