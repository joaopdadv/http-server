package main

import (
	"fmt"
	"log"
	"net/rpc"
)

const (
	connection = "tcp"
	host       = "localhost"
	port       = ":42069"
)

// Mesmos types do server
type SumArgs struct{ A, B int }
type SumReply struct{ Result int }

func main() {
	client, err := rpc.Dial(connection, host+port)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer client.Close()

	args := &SumArgs{A: 10, B: 32}
	var reply SumReply

	if err := client.Call("Math.Sum", args, &reply); err != nil {
		log.Fatal("call:", err)
	}

	fmt.Println("Resultado:", reply.Result) // 42
}
