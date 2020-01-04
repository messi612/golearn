package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/messi612/golearn/common"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
	}
	service := os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	common.CheckError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	common.CheckError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	common.CheckError(err)

	result, err := ioutil.ReadAll(conn)
	common.CheckError(err)

	fmt.Println(string(result))
	os.Exit(0)
}
