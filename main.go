package main

import (
	"log"
	"net"
)

func main() {
	testNet()
}

func testNet() {
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		log.Println(addr)
	}

	interfaces, _ := net.Interfaces()
	for _, interf := range interfaces {
		log.Println(interf)
	}
}
