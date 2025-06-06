package main

import (
	"fmt"
	"os"
	"strings"
)

/* display the number of lines, words, and bytes in a file */

func ws() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cat.go <filename>")
		return
	}

	fileName := os.Args[1]
	fileContent, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("error while reading the file content")
		return
	}
	byteToStringFormat := string(fileContent)
	individual_words := strings.Split(byteToStringFormat, " ")

	fmt.Println("Lines:", strings.Count(byteToStringFormat, "\n")+1)
	fmt.Println("Words:", len(individual_words))
	fmt.Println("Bytes:", len(byteToStringFormat))
}

func main() {
	ws()
}
