package thread_pool

import (
	"fmt"
	"testing"
	"time"
	"github.com/ziyouchutuwenwu/objective_go/thread/msg_thread"
)

var thread1 *msg_thread.Thread
var thread2 *msg_thread.Thread

type Test struct{}

func (test *Test) MyThreadCallBack1(thread *msg_thread.Thread, info interface{}) {
	fmt.Println("thread1")

	time.Sleep(1 * time.Second)

	info2 := "222"
	msgObject2 := new(msg_thread.MsgObject)
	msgObject2.Info = info2
	thread2.SendMsg(msgObject2)
}

func (test *Test) MyThreadCallBack2(thread *msg_thread.Thread, info interface{}) {
	fmt.Println("thread2")

	time.Sleep(1 * time.Second)

	info1 := "111"
	msgObject1 := new(msg_thread.MsgObject)
	msgObject1.Info = info1
	thread1.SendMsg(msgObject1)
}

func (test *Test) MyThreadMsgCallBack1(thread *msg_thread.Thread, infoObject interface{}) {
	fmt.Println("msgCallBack1", infoObject)

	info2 := "fuck from thread1 to thread2"
	msgObject2 := new(msg_thread.MsgObject)
	msgObject2.Info = info2
	thread2.SendMsg(msgObject2)
}

func (test *Test) MyThreadMsgCallBack2(thread *msg_thread.Thread, infoObject interface{}) {
	fmt.Println("msgCallBack2", infoObject)

	info1 := "fuck from thread2 to thread1"
	infoObject1 := new(msg_thread.MsgObject)
	infoObject1.Info = info1
	thread1.SendMsg(infoObject1)

	//可选是否关闭，一般不需要关闭，线程池关闭的时候，自动关闭此线程
	//thread1.Stop()
}

func (test *Test) MyThreadPoolCallBack(threadPool *ThreadPool) {
	fmt.Println("thread_pool")
	time.Sleep(2 * time.Second)

	threadPool.Stop()
}

func TestThreadPool(t *testing.T) {

	threadPool := Create()
	threadPool.Init()

	test := new(Test)

	thread1 = msg_thread.Create()
	thread1.Init()
	thread1.Tag = "111"
	thread1.SetThreadCallBack(test.MyThreadCallBack1)
	thread1.SetMsgCallBack(test.MyThreadMsgCallBack1)
	thread1.SetWaitMode(false)
	threadPool.Add(thread1)

	thread2 = msg_thread.Create()
	thread2.Init()
	thread2.Tag = "222"
	thread2.SetThreadCallBack(test.MyThreadCallBack2)
	thread2.SetMsgCallBack(test.MyThreadMsgCallBack2)
	thread2.SetWaitMode(false)
	threadPool.Add(thread2)

	threadPool.SetWaitMode(false)
	threadPool.SetCallBack(test.MyThreadPoolCallBack)
	threadPool.Start()

	threadPool.SetWaitMode(true)
	threadPool.Wait()
}
