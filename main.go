package main

import "fmt"

func main() {
	MyPrint("5")
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
