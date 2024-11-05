package test

import (
	"net"
)

func GetFreePort() (int, error) {
	port, err := FindFreePort("")
	return int(port), err
}

func FindFreePort(ip string) (uint32, error) {
	ln, err := net.Listen("tcp", ip+":0")
	if err != nil {
		return 0, err
	}
	if err := ln.Close(); err != nil {
		return 0, err
	}
	return uint32(ln.Addr().(*net.TCPAddr).Port), nil
}
