package dynamic_method

import (
	"fmt"
	"testing"
)

type YourT1 struct {
}

func (y *YourT1) MethodBar() {
	fmt.Println("MethodBar called")
}

type YourT2 struct {
}

func (y *YourT2) MethodFoo(i int, oo string) {
	fmt.Println("MethodFoo called", i, oo)
}

func TestMethod(t *testing.T) {
	InvokeObjectMethod(new(YourT2), "MethodFoo", 10, "abc")
	InvokeObjectMethod(new(YourT1), "MethodBar")
}