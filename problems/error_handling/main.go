package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	var maxLineNumber = 2

	file, err := os.OpenFile("notes.csv", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for range maxLineNumber {
		line, readErr := reader.ReadString('\n')
		if readErr != nil {
			if readErr == io.EOF {
				break // Reached end of file
			}
			fmt.Printf("Error reading line: %v\n", readErr)
			return
		}
		fmt.Print(line)
	}

}
