package main

import (
	"fmt"
	"github.com/messi612/golearn/common"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "USAGE: %s host:port", os.Args[0])
	}

	service := os.Args[1]

	conn, err := net.Dial("tcp", service)
	common.CheckError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	common.CheckError(err)

	result, err := ioutil.ReadAll(conn)
	common.CheckError(err)

	fmt.Println(string(result))
	os.Exit(0)
}
