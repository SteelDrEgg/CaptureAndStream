package util

import (
	"net"
	"os"
	"strings"
)

func GetIP() string {
	name, _ := os.Hostname()
	addr, _ := net.LookupHost(name)
	for i, _ := range addr {
		if strings.Contains(addr[i], ".") && addr[i] != "127.0.0.1" {
			return addr[i]
		}
	}
	return "127.0.0.1"
}
