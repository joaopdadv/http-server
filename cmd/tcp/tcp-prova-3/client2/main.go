package main

import (
	"fmt"
	"log"
	"net"
)

const (
	connection = "tcp"
	addr       = "localhost:7777"
)

func main() {

	// socket
	// connect
	// send
	// close

	conn, err := net.Dial(connection, addr)
	if err != nil {
		log.Fatalf("Error resolving tcp address: " + err.Error())
	}
	defer conn.Close()
	fmt.Println("Connected tcp address: " + conn.LocalAddr().Network())

	data := make([]byte, 256);
	n, err := conn.Read(data)

	if err != nil {
		log.Fatalf("Error resolving tcp address: " + err.Error())
	}

	fmt.Println("> ", string(data[:n]))

	fmt.Println("Connection to ", conn.RemoteAddr(), "closed")
}