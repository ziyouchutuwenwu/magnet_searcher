package basic_timer

import (
	"fmt"
	"time"
	"github.com/ziyouchutuwenwu/objective_go/exception"
)

type TimerCallBack func(timer *Timer, argObject interface{})

type Timer struct {
	duration                time.Duration

	timer                   *time.Ticker
	shouldDaemonChannelQuit chan bool
	isDaemonMode            bool

	timerCallBack           TimerCallBack

	Object                  interface{}
	Tag                     string
}

func Create() *Timer {
	timer := new(Timer)
	return timer
}

func (this *Timer) SetDuration(duration time.Duration) {
	this.duration = duration
}

func (this *Timer) SetTimerCallBack(callBackProc TimerCallBack) {
	this.timerCallBack = callBackProc
}

func (this *Timer) SetWaitMode(shouldWait bool) {
	this.isDaemonMode = shouldWait
}

func (this *Timer) Start() {

	go this.timerProc(this.duration)

	if this.isDaemonMode {
		this.daemonRoutine()
	}
}

func (this *Timer) Stop() {
	go this.notifyDaemonRoutineExit()

	this.timer.Stop()
}

func (this *Timer) notifyDaemonRoutineExit() {
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

func (this *Timer) Init() {
	this.Tag = ""
	this.duration = 1
	this.shouldDaemonChannelQuit = make(chan bool)
	this.isDaemonMode = false
	this.timer = nil
	this.timerCallBack = nil
}

func (this *Timer) ReInit() {
	this.Init()
}

func (this *Timer) timerProc(second time.Duration) {
	this.timer = time.NewTicker(second * time.Second)

	for {
		select {
		case <-this.timer.C:
			if nil != this.timerCallBack {
				this.timerCallBack(this, this.Object)
			}
		}
	}
}

func (this *Timer) daemonRoutine() {
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
