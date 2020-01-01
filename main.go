package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	select {
	case ch <- 1:
	case ch <- 2:
	case ch <- 3:
	}
	value1 := <-ch
	fmt.Println(value1)
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
