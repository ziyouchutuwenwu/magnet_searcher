package msg_timer

import (
	"fmt"
	"github.com/ziyouchutuwenwu/objective_go/exception"
	"time"
)

type MsgObject struct {
	cmd     string
	Info    interface{}
}

type TimerCallBack func(timer *Timer, argObject interface{})
type MsgCallBack func(timer *Timer, infoObject interface{})

type Timer struct {
	msgChannel              chan *MsgObject
	msgCallBack             MsgCallBack

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

func (this *Timer) SendMsg(msgObject *MsgObject) {
	if this.msgChannel != nil {
		go this.sendMsg(msgObject)
	}
}
func (this *Timer) sendMsg(msgObject *MsgObject) {
	exception.Try(
		func() {
			this.msgChannel <- msgObject
		},
		func(exception interface{}) {
			fmt.Println(exception)
		})
}

func (this *Timer) SetMsgCallBack(callBackProc MsgCallBack) {
	this.msgCallBack = callBackProc
}

func (this *Timer) SetTimerCallBack(callBackProc TimerCallBack) {
	this.timerCallBack = callBackProc
}

func (this *Timer) SetWaitMode(shouldWait bool) {
	this.isDaemonMode = shouldWait
}

func (this *Timer) Start() {

	if nil == this.timerCallBack && nil == this.msgCallBack {
		panic("timer callBacks not set")
	} else {
		go this.timerProc(this.duration)
	}

	go this.msgRoutine()

	if this.isDaemonMode {
		this.daemonRoutine()
	}
}

func (this *Timer) Restart() {
	this.Stop()
	this.Start()
}

func (this *Timer) Stop() {
	go this.notifyDaemonRoutineExit()

	this.timer.Stop()

	msgObject := new(MsgObject)
	msgObject.cmd = "quit"
	this.SendMsg(msgObject)
}

func (this *Timer) Init() {
	this.Tag = ""
	this.duration = 1
	this.timer = nil
	this.msgChannel = make(chan *MsgObject)
	this.shouldDaemonChannelQuit = make(chan bool)
	this.isDaemonMode = false
	this.timerCallBack = nil
	this.msgCallBack = nil
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

func (this *Timer) msgRoutine() {
	for {
		select {
		case msg, isChanOpen := <-this.msgChannel:
			if isChanOpen {
				if msg.cmd == "quit" {
					goto _exit
				} else {
					if nil == this.timerCallBack && nil == this.msgCallBack {
						panic("timer callBacks not set")
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
	fmt.Println("timer exit,tag ", this.Tag)
	close(this.msgChannel)
	this.msgChannel = nil
	return
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