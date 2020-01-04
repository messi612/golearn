package main

import "sync"
import "fmt"

var a string
var once sync.Once

func setup() {
	a = "hello world"
	fmt.Println("setup done")
}

func doprint(ch chan int) {
	once.Do(setup)
	fmt.Println(a)
	ch <- 1
}

func main() {
	chs := make([]chan int, 3)
	for i := 0; i < 3; i++ {
		chs[i] = make(chan int)
		go doprint(chs[i])
	}

	for _, ch := range chs {
		<-ch
	}
}
