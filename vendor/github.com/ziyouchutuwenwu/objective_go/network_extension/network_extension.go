package network_extension

import (
	"net"
	"strings"
)

func GetLocalIps() []string {
	var ips []string

	address, err := net.InterfaceAddrs()
	if nil != err {
		return ips
	}

	for _, addr := range address {
		if strings.Contains(addr.String(), ":") {
			continue
		}
		ip := strings.Split(addr.String(), "/")[0]
		ips = append(ips, ip)
	}

	return ips
}
