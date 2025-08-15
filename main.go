package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatalf("Erro abrindo arquivo: %v", err)
	}
	defer file.Close() // defer -> quando a função main terminar, o arquivo será fechado

	for {
		data := make([]byte, 8)
		_, err := file.Read(data)

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Erro lendo arquivo: %v", err)
		}

		fmt.Printf("read: %s\n", data)
	}
}
