package object

import "testing"

func TestObject(t *testing.T) {

	var a = Create()
	a.Value = 40

	var b = Create()
	b.Value = 100

	t.Log(a.GetValueTypeName())
	t.Log(b.GetValueTypeName())
	if a.GetValueType() == b.GetValueType() {
		t.Log("same type")
	}

	t.Log(a.IsEqualToObject(b))
}
