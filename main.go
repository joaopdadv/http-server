package main

import (
	"errors"
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

	// como se fosse um for await no js - cada out <- valor dispara esse range
	for i := range getLinesChannel(file) {
		fmt.Printf("read: %s\n", i)
	}
}

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
					fmt.Printf("read: %s\n", currentLine)
				}
				if errors.Is(err, io.EOF) {
					break
				}
				log.Fatalf("Erro lendo arquivo: %v", err)
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
