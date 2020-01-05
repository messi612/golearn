package main

import "net/rpc"

import "log"

import "fmt"

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")

	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &Args{3, 5}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d * %d = %d \n", args.A, args.B, reply)

	quotient := new(Quotient)
	divCall := client.Go("Arith.Divide", args, &quotient, nil)
	<-divCall.Done
	fmt.Printf("quo = %d,rem = %d \n", quotient.Quo, quotient.Rem)
}
