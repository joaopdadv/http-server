package main

import (
	"fmt"
	"log"
	"net"
)

// server -> c1 -> c2

// c1 manda para server e server repassa para c2

const (
	connection = "tcp"
	port       = ":7777"
)

// Server
func main() {

	// socket
	// bind
	// accept 1
	// listen 1
	// accept 2
	// send 2

	listener, err := net.Listen(connection, port)
	if err != nil {
		log.Fatal("Error listening " + connection + " traffic at " + port + ":" + err.Error())
	}
	defer listener.Close()

	fmt.Println("Listening for TCP traffic on", port)

	conn, errAccept := listener.Accept() // aceita client 1
	if errAccept != nil {
		log.Fatal("Error accepting connection 1 from " + connection + " port " + port)
	}
	defer conn.Close()
	
	fmt.Println("Accepted connection 1 from", conn.RemoteAddr())

	data := make([]byte, 256)

	n, err := conn.Read(data);

	if err != nil {
		log.Fatalf("Error reading bytes: %v", err)
	}

	conn.Close();
	fmt.Println("Connection to ", conn.RemoteAddr(), "closed")

	conn, errAccept = listener.Accept() // aceita client 2
	if errAccept != nil {
		log.Fatal("Error accepting connection 2 from " + connection + " port " + port)
	}
	defer conn.Close()

	fmt.Println("Accepted connection 2 from", conn.RemoteAddr())

	conn.Write(data[:n])
	
	conn.Close();
	fmt.Println("Connection to ", conn.RemoteAddr(), "closed")
}