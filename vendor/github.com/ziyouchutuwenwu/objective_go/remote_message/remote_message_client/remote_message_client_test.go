package remote_message_client

import "testing"

func TestRemoteMessageClient(t *testing.T) {
	remoteMessageClient := Create()
	remoteMessageClient.Init()

	remoteMessageClient.Ip = "0.0.0.0"
	remoteMessageClient.Port = 12345

	param := make(map[string]string)
	param["aa"] = "bb"
	param["cc"] = "dd"
	param["ee"] = "ff"
	result := remoteMessageClient.SendMessage(1111, param)
	t.Log(result)
}
