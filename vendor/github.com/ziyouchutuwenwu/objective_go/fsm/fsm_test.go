package fsm

import (
	"testing"
	"log"
	"errors"
	"time"
)

type MyObject struct {
}

func (myObject *MyObject) State1(event string, param interface{}, timeout int) (nextState string, timeOut int, err error) {
	log.Println("in State1 ", event, param.(int))
	return "State2", 0, nil //如果timeout大于0，则在timeout毫秒后，自动调用下一个状态,下一个状态的event为timeout
}

func (myObject *MyObject) State2(event string, param interface{}, timeout int) (nextState string, timeOut int, err error) {
	log.Println("in State2 ", event, param.(int))
	return "stop", 0, errors.New("stop ok") //nextstate=stop则停止状态机，err为停止原因
}

/*
参数：
	event是事件名
	param为事件的参数
	timeout > 0 表示这是一个延时事件

返回值：
	nextState，必须和状态回调函数同名，如果为"stop"则表示没有后续的状态，状态机停止。
	timeOut > 0 表示延时回调，将在timeout时间后，产生一个timeout事件。
*/
func myFsm() {
	myObjectFsm := NewFSM(&MyObject{})
	myObjectFsm.Init("State1")

	myObjectFsm.SendEvent("Do", 1)
	myObjectFsm.SendEvent("Do", 1)
	time.Sleep(time.Second * 2)
}

func TestFSM(t *testing.T){
	myFsm()
}