package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	grep()
}

func grep() {

	if len(os.Args[1]) < 2 {
		fmt.Println("Something went while reading the file")
		return
	}

	fileName := os.Args[1]
	target := os.Args[2]

	file, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Something went while reading the content of file")
	}

	content := string(file)

	words := strings.Split(content, " ")

	for i := 0; i < len(words); i++ {
		if words[i] == target {
			highlighted := strings.Replace(content, words[i], "\033[31m"+words[i]+"\033[0m", i)
			fmt.Println(highlighted)
		}
	}
}
