package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const (
	connection = "udp"
)

func getInfoChannel(f io.ReadCloser) <-chan string {

	out := make(chan string)

	go func() {
		defer close(out)

		for {
			data := make([]byte, 2048) // UDP -> cada read corresponde a um datagrama inteiro
			n, err := f.Read(data)

			if err != nil {
				break
			}

			out <- string(data[:n])
		}
	}()
	return out
}

func main() {

	args := os.Args

	if len(args) != 2 {
		log.Fatal("Error: command should be 'go run main.go <port>'")
	}

	port := "localhost:" + os.Args[1]

	addr, err := net.ResolveUDPAddr(connection, port)
	if err != nil {
		log.Fatal(err.Error())
	}

	conn, err := net.ListenUDP(connection, addr)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()

	fmt.Println("Listening for UDP traffic on", port)

	lines := getInfoChannel(conn)
	for line := range lines {
		fmt.Printf("%s", line)
	}

}
