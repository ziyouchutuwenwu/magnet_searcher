package thread_pool

import (
	"github.com/ziyouchutuwenwu/objective_go/data_foundation/data_type/array"
	"github.com/ziyouchutuwenwu/objective_go/thread"
	"github.com/ziyouchutuwenwu/objective_go/exception"
	"fmt"
)

type CallBack func(threadPool *ThreadPool)

type ThreadPool struct {
	threadArray             array.Array

	shouldDaemonChannelQuit chan bool
	isDaemonMode            bool
	threadPoolCallBack      CallBack
}

func Create() *ThreadPool {
	threadPool := new(ThreadPool)
	return threadPool
}

func (this *ThreadPool) Init() {
	this.threadArray = array.Create()
	this.shouldDaemonChannelQuit = make(chan bool)
	this.isDaemonMode = false
	this.threadPoolCallBack = nil
}

func (this *ThreadPool) ReInit() {
	this.Init()
}

func (this *ThreadPool) SetCallBack(callBackProc CallBack) {
	this.threadPoolCallBack = callBackProc
}

func (this *ThreadPool) SetWaitMode(shouldWait bool) {
	this.isDaemonMode = shouldWait
}

func (this *ThreadPool) Add(thread thread.IThread) {
	this.threadArray.AddObject(thread)
}

func (this *ThreadPool) Remove(thread thread.IThread) {
	index := this.threadArray.FindObject(thread)
	if index >= 0 {
		this.threadArray.RemoveObjectAtIndex(index)
	}
}

func (this *ThreadPool) Start() {
	for i := 0; i < this.threadArray.Len(); i++ {
		thread := this.threadArray.GetObjectAtIndex(i).(thread.IThread)
		thread.Start()
	}

	if nil == this.threadPoolCallBack {
		panic("threadPool callback not set")
	} else {
		if nil != this.threadPoolCallBack {
			go this.threadPoolCallBack(this)
		}
	}

	if this.isDaemonMode {
		this.daemonRoutine()
	}
}

func (this *ThreadPool) Wait() {
	if this.isDaemonMode {
		this.daemonRoutine()
	}
}

func (this *ThreadPool) Stop() {
	for i := 0; i < this.threadArray.Len(); i++ {
		thread := this.threadArray.GetObjectAtIndex(i).(thread.IThread)
		thread.Stop()
	}
	this.threadArray.RemoveAllObjects()

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

func (this *ThreadPool) daemonRoutine() {
	for {
		select {
		case shouldDaemonQuit, isChanOpen := <-this.shouldDaemonChannelQuit:
			if isChanOpen {
				if shouldDaemonQuit {
					fmt.Println("pool will exit")
					close(this.shouldDaemonChannelQuit)
					return
				}
			}
		}
	}
}
