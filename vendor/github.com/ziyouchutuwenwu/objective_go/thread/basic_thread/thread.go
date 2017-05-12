package basic_thread

import (
	"fmt"
	"github.com/ziyouchutuwenwu/objective_go/exception"
)

type ThreadCallBack func(thread *Thread, argObject interface{})

type Thread struct {
	shouldDaemonChannelQuit chan bool
	isDaemonMode    bool
	threadCallBack  ThreadCallBack

	Object          interface{}
	Tag             string
}

func Create() *Thread {
	thread := new(Thread)
	return thread
}

func (this *Thread) Start() {
	go this.workingRoutine()

	if this.isDaemonMode {
		this.daemonRoutine()
	}
}

func (this *Thread) Stop() {
	go this.notifyDaemonRoutineExit()
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

func (this *Thread) SetCallBack(callBackProc ThreadCallBack) {
	this.threadCallBack = callBackProc
}

func (this *Thread) SetWaitMode(shouldWait bool) {
	this.isDaemonMode = shouldWait
}

func (this *Thread) Init() {
	this.Tag = ""
	this.shouldDaemonChannelQuit = make(chan bool)
	this.isDaemonMode = false
	this.threadCallBack = nil
}

func (this *Thread) workingRoutine() {
	if nil == this.threadCallBack {
		panic("basic_thread callBack not set")
	}
	this.threadCallBack(this, this.Object)
	if this.isDaemonMode {
		this.Stop()
	}
}

func (this *Thread) daemonRoutine() {
	for {
		select {
		case shouldDaemonQuit, isChanOpen := <-this.shouldDaemonChannelQuit:
			if isChanOpen {
				if shouldDaemonQuit {
					close(this.shouldDaemonChannelQuit)
					return
				}
			}
		}
	}
}
