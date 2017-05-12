package basic_thread

import (
	"fmt"
	"testing"
	"time"
)

type Test struct{}

func (test *Test) MyThreadCallBack(thread *Thread, argObject interface{}) {

	fmt.Println("thread work begin")
	time.Sleep(2 * time.Second)
	fmt.Println("thread work done", argObject)
}

func TestThread(t *testing.T) {

	test := new(Test)

	thread := Create()
	thread.Init()

	thread.Tag = "111"
	thread.Object = "asd"

	thread.SetCallBack(test.MyThreadCallBack)
	thread.SetWaitMode(true)
	thread.Start()
}