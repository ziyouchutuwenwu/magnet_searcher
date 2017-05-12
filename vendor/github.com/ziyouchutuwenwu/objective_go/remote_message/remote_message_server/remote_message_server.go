package remote_message_server

import (
	"github.com/hoisie/web"
	"github.com/ziyouchutuwenwu/objective_go/data_foundation/data_convert"
)

type CallBack func(cmd string, params map[string]string) string

type RemoteMessageServer struct {
	Ip   string
	Port int

	msgCallBack CallBack
}

func Create() *RemoteMessageServer {
	object := new(RemoteMessageServer)
	return object
}

func (this *RemoteMessageServer) Init() {
	this.Ip = ""
	this.Port = 0
}

func (this *RemoteMessageServer) SetCallBack(callBackProc CallBack) {
	this.msgCallBack = callBackProc
}

func (this *RemoteMessageServer) remoteMessageProc(ctx *web.Context, val string) string {

	var cmd string
	params := make(map[string]string)

	count := 0
	for key, value := range ctx.Params {

		if 0 == count {
			if key == "cmd" {
				cmd = value
			}
		} else {
			params[key] = value
		}
		count++
	}

	return this.msgCallBack(cmd, params)
}

func (this *RemoteMessageServer) Run() {
	web.Get("/(.*)", this.remoteMessageProc)

	serverStr := this.Ip + ":" + data_convert.IntToStr(this.Port)
	web.Run(serverStr)
}

func (this *RemoteMessageServer) Stop() {
	web.Close()
}
