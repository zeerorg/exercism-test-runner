package main

import "net"

func GetListener() net.Listener {
	l, err := net.Listen("unix", "/tmp/exercism_test_socket")
	if err != nil {
		panic(err)
	}
	return l
}
