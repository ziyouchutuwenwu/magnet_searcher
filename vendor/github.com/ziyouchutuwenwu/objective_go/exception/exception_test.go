package exception

import (
	"fmt"
	"testing"
)

func say(s string) {
	fmt.Println(s)
}

func TestExcept(t *testing.T) {
	say("Hello")
	Try(
		func() {
			panic("World")
		},
		func(exp interface{}) {
			t.Log("catch", exp)
		})
	say("end")
}