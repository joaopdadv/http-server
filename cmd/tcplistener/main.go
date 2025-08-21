package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

const (
	connection = "tcp"
	port       = ":42069"
)

func getLinesChannel(f io.ReadCloser) <-chan string {

	out := make(chan string)

	// roda go routines em paralelo para processamento (mais rápido que threads nativas)
	go func() {
		defer close(out)
		currentLine := ""

		for {
			data := make([]byte, 8)
			n, err := f.Read(data)

			if err != nil {
				if currentLine != "" {
					out <- currentLine
				}
				if errors.Is(err, io.EOF) {
					break
				}
				log.Fatalf("Error reading bytes: %v", err)
				break
			}

			lineStrings := strings.Split(string(data[:n]), "\n") // data[:n] é importante para não vazar bytes nulos ou lixo na string()

			// Se tem apenas 1 item no vetor lineStrings, len(lineStrings)-1 = 0, ou seja, loop não roda pois 0 < 0 = false
			for i := 0; i < len(lineStrings)-1; i++ {
				out <- currentLine + lineStrings[i]
				currentLine = ""
			}

			// coloca o último elemento sempre em currentLine
			currentLine += lineStrings[len(lineStrings)-1]
		}
	}()
	return out
}

func main() {

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

		lines := getLinesChannel(conn)
		for line := range lines {
			fmt.Printf("%s\n", line)
		}

		fmt.Println("Connection to ", conn.RemoteAddr(), "closed")
	}
}
