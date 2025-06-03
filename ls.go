package main

import (
	"fmt"
	"os"
)

func ls() {
	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	files, err := os.ReadDir(currentWorkingDirectory)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		info, err := os.Stat(file.Name())
		if err != nil {
			fmt.Println(err)
		}

		mode := info.Mode()
		name := file.Name()

		if mode.IsRegular() && (mode&0111 != 0) {
			fmt.Printf("\033[1;32m%s\033[0m", name)
		} else {
			fmt.Print("  ", name)
		}
	}
	fmt.Println()
}
