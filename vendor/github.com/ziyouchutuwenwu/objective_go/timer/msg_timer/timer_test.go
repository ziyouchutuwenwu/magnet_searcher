package msg_timer

import (
	"fmt"
	"testing"
	"time"
)

var timer1 *Timer
var timer2 *Timer

type Test struct {
	count int
}

func (test *Test) MyTimerCallBack(timer *Timer, argObject interface{}) {

	test.count++

	fmt.Println("timerCallBack", timer.Tag, argObject)

	if timer1 == timer {

		msgObject := new(MsgObject)
		msgObject.Info = "got msg from 1"
		timer2.SendMsg(msgObject)
	}

	if test.count >= 4 {

		timer1.Restart()

		time.Sleep(1)

		timer2.Stop()
	}
}

func (test *Test) MyTimerMsgCallBack(timer *Timer, infoObject interface{}) {
	fmt.Println("msgCallBack", timer.Tag, infoObject)

	if timer == timer2 {
		msgObject := new(MsgObject)
		msgObject.Info = "got msg from 2"
		timer1.SendMsg(msgObject)
	}
}

func TestTimer(t *testing.T) {

	test := new(Test)

	timer1 = Create()
	timer1.Init()
	timer1.Tag = "111"
	timer1.SetMsgCallBack(test.MyTimerMsgCallBack)
	timer1.SetDuration(3)
	timer1.SetTimerCallBack(test.MyTimerCallBack)
	//这里为false
	timer1.SetWaitMode(false)
	timer1.Object = "111 args"
	timer1.Start()

	timer2 = Create()
	timer2.Init()
	timer2.Tag = "222"
	timer2.SetMsgCallBack(test.MyTimerMsgCallBack)
	timer2.SetDuration(5)
	timer2.SetTimerCallBack(test.MyTimerCallBack)
	//这里为true
	timer2.SetWaitMode(true)

	timer2.Object = "222 args"
	timer2.Start()
}
