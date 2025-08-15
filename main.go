package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatalf("Erro abrindo arquivo: %v", err)
	}
	defer file.Close() // defer -> quando a função main terminar, o arquivo será fechado

	currentLine := ""

	for {
		data := make([]byte, 8)
		_, err := file.Read(data)

		if err == io.EOF {
			fmt.Printf("read: %s\n", currentLine)
			break
		}
		if err != nil {
			log.Fatalf("Erro lendo arquivo: %v", err)
		}

		lineStrings := strings.Split(string(data), "\n")

		currentLine += lineStrings[0]

		if len(lineStrings) > 1 {
			fmt.Printf("read: %s\n", currentLine)

			currentLine = lineStrings[1]
		}
	}
}
