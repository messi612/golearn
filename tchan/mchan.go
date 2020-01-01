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
}
