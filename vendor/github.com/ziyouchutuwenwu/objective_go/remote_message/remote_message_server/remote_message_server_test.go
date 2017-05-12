package remote_message_server

import (
	"fmt"
	"testing"
)

func MsgCallBack(cmd string, params map[string]string) string {
	fmt.Println(cmd, params)
	//做测试,返回客户端数据
	return cmd
}

func TestRemoteMessageServer(t *testing.T) {
	remoteMessageServer := Create()
	remoteMessageServer.Init()

	remoteMessageServer.Ip = "0.0.0.0"
	remoteMessageServer.Port = 12345
	remoteMessageServer.SetCallBack(MsgCallBack)
	remoteMessageServer.Run()
}
