package basic_timer

import (
	"fmt"
	"testing"
)

type Test struct {
	count int
}

func (test *Test) MyTimerCallBack(timer *Timer, argObject interface{}) {

	test.count++

	fmt.Println("timerCallBack", timer.Tag, argObject)

	if test.count >= 4 {
		timer.Stop()
	}
}

func TesTimer(t *testing.T) {T
	test := new(Test)
	timer := Create()
	timer.Init()
	timer.Tag = "111"
	timer.SetDuration(1)
	timer.SetTimerCallBack(test.MyTimerCallBack)
	timer.SetWaitMode(true)

	timer.Object = "my timer info"

	timer.Start()
}