package msg_thread

import (
	"fmt"
	"github.com/ziyouchutuwenwu/objective_go/exception"
)

type MsgObject struct {
	cmd     string
	Info    interface{}
}

type ThreadCallBack func(thread *Thread, argObject interface{})
type MsgCallBack func(thread *Thread, infoObject interface{})

type Thread struct {
	msgChannel              chan *MsgObject
	msgCallBack             MsgCallBack

	shouldDaemonChannelQuit chan bool
	isDaemonMode            bool
	threadCallBack          ThreadCallBack

	Object                  interface{}
	Tag                     string
}

func Create() *Thread {
	thread := new(Thread)
	return thread
}

func (this *Thread) Start() {

	if nil == this.threadCallBack && nil == this.msgCallBack {
		panic("workingThread callBacks not set")
	} else {
		if nil != this.threadCallBack {
			go this.threadCallBack(this, this.Object)
		}
	}

	go this.workingMsgRoutine()

	if this.isDaemonMode {
		this.daemonRoutine()
	}
}

func (this *Thread) Stop() {

	go this.notifyDaemonRoutineExit()

	msgObject := new(MsgObject)
	msgObject.cmd = "quit"
	this.SendMsg(msgObject)
}

func (this *Thread) notifyDaemonRoutineExit() {
	if this.isDaemonMode {
		exception.Try(
			func() {
				this.shouldDaemonChannelQuit <- true
			},
			func(exception interface{}) {
				fmt.Println(exception)
			})
	}
}

func (this *Thread) SendMsg(msgObject *MsgObject) {
	if this.msgChannel != nil {
		go this.sendMsg(msgObject)
	}
}

func (this *Thread) sendMsg(msgObject *MsgObject) {
	exception.Try(
		func() {
			this.msgChannel <- msgObject
		},
		func(exception interface{}) {
			fmt.Println(exception)
		})
}

func (this *Thread) SetMsgCallBack(callBackProc MsgCallBack) {
	this.msgCallBack = callBackProc
}

func (this *Thread) SetThreadCallBack(callBackProc ThreadCallBack) {
	this.threadCallBack = callBackProc
}

func (this *Thread) SetWaitMode(shouldWait bool) {
	this.isDaemonMode = shouldWait
}

func (this *Thread) Init() {
	this.Tag = ""
	this.msgChannel = make(chan *MsgObject)
	this.shouldDaemonChannelQuit = make(chan bool)
	this.isDaemonMode = false
	this.msgCallBack = nil
	this.threadCallBack = nil
}

func (this *Thread) workingMsgRoutine() {
	for {
		select {
		case msg, isChanOpen := <-this.msgChannel:
			if isChanOpen {
				if msg.cmd == "quit" {
					goto _exit
				} else {
					if nil == this.threadCallBack && nil == this.msgCallBack {
						panic("workingThread callBacks not set")
					} else {
						if nil != this.msgCallBack {
							this.msgCallBack(this, msg.Info)
						}
					}
				}
			}
		}
	}

	_exit:
	fmt.Println("thread exit,tag ", this.Tag)
	close(this.msgChannel)
	this.msgChannel = nil
	return
}

func (this *Thread) daemonRoutine() {
	for {
		select {
		case shouldDaemonQuit, isChanOpen := <-this.shouldDaemonChannelQuit:
			if isChanOpen {
				if shouldDaemonQuit {
					fmt.Println("thread quit in daemon")
					close(this.shouldDaemonChannelQuit)
					return
				}
			}
		}
	}
}