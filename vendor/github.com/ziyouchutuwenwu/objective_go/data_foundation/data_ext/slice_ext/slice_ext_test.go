package slice_ext

import (
	"testing"
	"sort"
	"fmt"
)

type MyObject struct {
	Name string
}

func TestSliceExt(t *testing.T) {
	var objects []*MyObject
	for i := 0; i < 8; i++ {
		object := new(MyObject)
		objects = append(objects, object)
	}
	t.Log(objects)
	leftObjects := SliceRemoveAtIndex(objects, 1).([]*MyObject)
	t.Log(leftObjects)

	leftObjects = SliceRemoveWithRange(objects, 1, 2).([]*MyObject)
	t.Log(leftObjects)

	t.Log(IsSliceEqual(objects, leftObjects))
}

func TestDeDuplicate(t *testing.T){
	b := []string{"a", "b", "c", "c", "e", "f", "a", "g", "b", "b", "c"}
	sort.Strings(b)
	fmt.Println(DeDuplicate(b))


	c := []int{1, 1, 2, 4, 6, 7, 8, 4, 3, 2, 5, 6, 6, 8}

	sort.Ints(c)
	fmt.Println(DeDuplicate(c))
}