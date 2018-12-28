package main

import "net"

var Listener, _ = net.Listen("unix", "/tmp/exercism_test_runner")
