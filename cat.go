// if we want to print the content of the file, we first of all needs to read the file

/*
steps 1: read a file
steps 2: save the content of the file somewhere
steps 3: display the stored content in the terminal in prettier format(will figured out later after implementing all the functionalities)
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cat.go <filename>")
		return
	}
	
	fileName := os.Args[1]

	content, err := os.ReadFile(fileName)
 	if err != nil {
		fmt.Println("Error while reading the file", err);
		return
	}

	fmt.Println(string(content))
}
