package network_extension

import "testing"

func TestNetworkExtension(t *testing.T) {
	ips := GetLocalIps()
	t.Log(ips)
}
