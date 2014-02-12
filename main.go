package main

import (
	// "bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	// "strings"
)

// func main() {
// 	//testNet()
// 	testDumpRequest()
// }

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
func testDumpRequest() {
	// s := "POST /v2/api/?login HTTP/1.1\r\nHost: www.baidu.com\r\nConnection: keep-alive\r\nContent-Length: 21\r\nUser-Agent: Mozilla/5.0 (Windows NT 5.1)\r\nContent-Type: application/x-www-form-urlencoded\r\n\r\nkey1=name1&key2=name2"
	// req, err := http.ReadRequest(bufio.NewReader(strings.NewReader(s)))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	req, _ := http.NewRequest("GET", "http://www.baidu.com", nil)
	req.ContentLength = 100
	// defer req.Body.Close()
	dump, err := httputil.DumpRequestOut(req, true)
	fmt.Println(string(dump))
	if err != nil {
		fmt.Println(err)
		return
	}
}
func init() {
	log.Println("init1")
}
func init() {
	log.Println("init2")
}
