package main

import "net"

var Listener, _ = net.Listen("tcp", ":8000")
