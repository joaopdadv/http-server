package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

const (
	connection = "tcp"
	port       = ":42069"
)

type SumArgs struct {
	A, B int
}
type SumReply struct {
	Result int
}

type Math int

func (m *Math) Sum(args *SumArgs, reply *SumReply) error {
	reply.Result = args.A + args.B
	return nil
}

func main() {

	if err := rpc.Register(new(Math)); err != nil {
		log.Fatal("erro ao registrar service:", err)
	}

	listener, err := net.Listen(connection, port)
	if err != nil {
		log.Fatal("Error listening " + connection + " traffic at " + port + ":" + err.Error())
	}
	defer listener.Close()
	fmt.Println("Listening for TCP traffic on", port)

	for {
		conn, errAccept := listener.Accept()

		if errAccept != nil {
			log.Fatal("Error accepting connection from " + connection + " port " + port)
		}

		fmt.Println("Accepted connection from", conn.RemoteAddr())

		go rpc.ServeConn(conn) // goroutine
	}
}
