package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	address = "localhost:42069"
	network = "udp"
)

func main() {

	addr, err := net.ResolveUDPAddr(network, address)
	if err != nil {
		log.Fatal("Error resolving udp address: " + err.Error())
	}

	conn, err := net.DialUDP(network, nil, addr)
	if err != nil {
		log.Fatal("Error resolving udp address: " + err.Error())
	}
	defer conn.Close()
	fmt.Println("Connected udp address: " + conn.LocalAddr().Network())

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, errInput := reader.ReadString('\n') // reads until newLine is found

		if errInput != nil {
			log.Fatal("Error reading line: " + errInput.Error())
			break
		}

		_, err := conn.Write([]byte(input))

		if err != nil {
			log.Fatal("Error resolving udp address: " + err.Error())
			break
		}
	}
}
