package main

import "net"

func GetListener() net.Listener {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	return l
}
