/*
抄自http://blog.csdn.net/sll1983/article/details/37560215
感谢作者sll1983
*/
package fsm

import (
	"errors"
	"reflect"
	"sync"
	"time"
	"unicode"
	"unicode/utf8"
	"log"
)

var typeOfError = reflect.TypeOf((*error)(nil)).Elem()

type FSM struct {
	sync.Mutex
	StopReason string
	rcvr       reflect.Value // receiver of methods for the service
	typ        reflect.Type  // type of the receiver
	method     map[string]reflect.Method
	event      chan Event
	quit       chan int
	state      string
	stopped    bool
}

type Event struct {
	event   string
	param   interface{}
	timeout int
}

func (fsm *FSM) IsStopped() bool {
	fsm.Lock()
	defer fsm.Unlock()
	return fsm.stopped
}

func (fsm *FSM) SendEvent(event string, param interface{}) {
	fsm.Lock()
	defer fsm.Unlock()
	if fsm.stopped {
		return
	}

	fsm.event <- Event{event, param, 0}

}

func (fsm *FSM) Init(start string) error {
	if _, ok := fsm.method[start]; !ok {
		return errors.New("not found state")
	}
	fsm.state = start
	go func() {
		for {
			select {
			case e := <-fsm.event:
				go fsm.CallState(e)
			case <-fsm.quit:
				goto close
			}
		}
	close:
		close(fsm.event)
		close(fsm.quit)
	}()

	return nil
}

func (fsm *FSM) CallState(e Event) {
	fsm.Lock()
	defer fsm.Unlock()
	if function, ok := fsm.method[fsm.state]; ok {
		returnValues := function.Func.Call([]reflect.Value{fsm.rcvr, reflect.ValueOf(e.event), reflect.ValueOf(e.param), reflect.ValueOf(e.timeout)})
		nextstate := returnValues[0].String()
		timeout := returnValues[1].Int()
		errInter := returnValues[2].Interface()
		errmsg := ""

		if errInter != nil {
			errmsg = errInter.(error).Error()
		}

		if nextstate == "stop" {
			fsm.Stop(errmsg)
			fsm.quit <- 1
			return
		}

		if errmsg != "" {
			//log.LogError(errmsg)
			log.Fatal(errmsg)
		}

		fsm.state = nextstate

		if timeout > 0 {
			go fsm.DelayCall(time.Duration(timeout))
		}
	}
}

func (fsm *FSM) DelayCall(timeout time.Duration) {
	select {
	case <-time.After(timeout * time.Millisecond):
		fsm.event <- Event{"timeout", 0, int(timeout)}
	}
}

func (fsm *FSM) Stop(message string) {
	fsm.StopReason = message
	fsm.stopped = true
}

func (fsm *FSM) Close() {
	fsm.Lock()
	defer fsm.Unlock()
	if fsm.stopped {
		return
	}

	fsm.quit <- 1
}

func NewFSM(fsm interface{}) *FSM {
	f := &FSM{typ: reflect.TypeOf(fsm), rcvr: reflect.ValueOf(fsm), event: make(chan Event), quit: make(chan int)}
	f.method = suitableMethods(f.typ, true)
	return f
}

func isExported(name string) bool {
	rune, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(rune)
}

func suitableMethods(typ reflect.Type, reportErr bool) map[string]reflect.Method {
	methods := make(map[string]reflect.Method)
	for m := 0; m < typ.NumMethod(); m++ {
		method := typ.Method(m)
		mtype := method.Type
		mname := method.Name

		if !isExported(mname) {
			continue
		}

		// Method needs four ins: receiver, string, interface{}, int.
		if mtype.NumIn() != 4 {
			if reportErr {
				log.Fatal("method", mname, "has wrong number of ins:", mtype.NumIn())
			}
			continue
		}

		// First arg must be a string.
		if mtype.In(1).Kind() != reflect.String {
			if reportErr {
				log.Fatal("method", mname, "arg1 type not a string:", mtype.In(1).Kind())
			}
			continue
		}
		// Second arg must be a interface.
		if mtype.In(2).Kind() != reflect.Interface {
			if reportErr {
				log.Fatal("method", mname, "arg2 type not a interface:", mtype.In(2).Kind())
			}
			continue
		}

		// Third arg must be a int.
		if mtype.In(3).Kind() != reflect.Int {
			if reportErr {
				log.Fatal("method", mname, "arg3 type not a int:", mtype.In(3).Kind())
			}
			continue
		}

		// Method needs three out.
		if mtype.NumOut() != 3 {
			if reportErr {
				log.Fatal("method", mname, "has wrong number of outs:", mtype.NumOut())
			}
			continue
		}

		if mtype.Out(0).Kind() != reflect.String {
			if reportErr {
				log.Fatal("method", mname, "out1 type not a string:", mtype.Out(0).Kind())
			}
			continue
		}
		if mtype.Out(1).Kind() != reflect.Int {
			if reportErr {
				log.Fatal("method", mname, "out1 type not a int:", mtype.Out(1).Kind())
			}
			continue
		}
		if mtype.Out(2) != typeOfError {
			if reportErr {
				log.Fatal("method", mname, "out3 type not a error:", mtype.In(2).Kind())
			}
			continue
		}
		methods[mname] = method
	}
	return methods
}