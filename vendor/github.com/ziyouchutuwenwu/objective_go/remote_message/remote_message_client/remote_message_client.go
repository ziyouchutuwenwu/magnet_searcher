package remote_message_client

import (
	"github.com/ziyouchutuwenwu/objective_go/data_foundation/data_convert"
	"io/ioutil"
	"net/http"
)

type RemoteMessageClient struct {
	Ip   string
	Port int
}

func Create() *RemoteMessageClient {
	object := new(RemoteMessageClient)
	return object
}

func (this *RemoteMessageClient) Init() {
	this.Ip = ""
	this.Port = 0
}

func (this *RemoteMessageClient) SendMessage(cmd int, params map[string]string) string {
	url := "http://" + this.Ip + ":" + data_convert.IntToStr(this.Port) + "/?" + "cmd=" + data_convert.IntToStr(cmd)
	for key, value := range params {
		paramInfo := "&" + key + "=" + value
		url += paramInfo
	}
	response, err := http.Get(url)
	if nil == err {
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		return string(body)
	}
	return "error"
}
