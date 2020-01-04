package main

import "fmt"

import "net"

import "github.com/messi612/golearn/common"

import "net/http"

import "log"

import "html"

func main() {
	MyPrint("5")

	fmt.Println(net.ParseIP("127.0.0.1"))
	ips, err := net.LookupIP("wwww.baidu.com")
	common.CheckError(err)
	fmt.Println(ips)
	fmt.Println(ips[0].DefaultMask())

	addr, err := net.ResolveIPAddr("ip4", "www.baidu.com")
	common.CheckError(err)
	fmt.Println(addr)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func MyPrint(args ...interface{}) {
	for idx, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value :", idx)
		case string:
			fmt.Println(arg, "is an string value:", idx)
		default:
			fmt.Println(arg, "is an unknown type:", idx)
		}
	}
}
