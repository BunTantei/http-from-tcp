package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const inputFilePath = "message.txt"

func main() {
	f, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("could not open %s: %s\n", inputFilePath, err)
	}
	defer f.Close()

	var currentLine string

	fmt.Printf("Reading data from %s\n", inputFilePath)
	fmt.Println("=====================================")

	for {
		b := make([]byte, 8)
		n, err := f.Read(b)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Printf("error: %s\n", err.Error())
			break
		}
		str := string(b[:n])
		currentLine += str
		if strings.Contains(currentLine, "\n") {
			parts := strings.Split(currentLine, "\n")
			for i := 0; i < len(parts)-1; i++ {
				fmt.Printf("read: %s\n", parts[i])
			}
			currentLine = parts[len(parts)-1]

		}
	}
	if currentLine != "" {
		fmt.Printf("read: %s\n", currentLine)
	}
}
