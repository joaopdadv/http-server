package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func getLinesChannel() <-chan string {
	out := make(chan string)

	connection := "tcp"
	port := ":42069"

	listener, err := net.Listen(connection, port)

	if err != nil {
		log.Fatal("Erro ao abrir conexão " + connection + " na porta " + port + ":" + err.Error())
	}
	defer listener.Close()

	conn, errAccept := listener.Accept()

	if errAccept != nil {
		log.Fatal("Erro ao aceitar conexão " + connection + " na porta " + port)
	}

	fmt.Println("Conexão " + connection + " na porta " + port + " foi aceita.")

	// roda go routines em paralelo para processamento (mais rápido que threads nativas)
	go func() {
		defer close(out)

		currentLine := ""

		for {
			data := make([]byte, 8)

			n, errCon := conn.Read(data)

			if errCon != nil {
				if currentLine != "" {
					fmt.Printf("%s\n", currentLine)
				}
				if errors.Is(errCon, io.EOF) {
					break
				}
				log.Fatalf("Erro lendo bytes: %v", errCon)
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
	// como se fosse um for await no js - cada out <- valor dispara esse range
	lines := getLinesChannel()
	for line := range lines {
		fmt.Printf("%s\n", line)
	}
}
