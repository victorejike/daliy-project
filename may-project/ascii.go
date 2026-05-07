package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run. [string]")
		return
	}

	input := os.Args[1]
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error : the file can't be found", err)
		return
	}
	line := strings.Split(string(file), "\n")

	for row := 0; row < 8; row++ {
		for _, ch := range input {
			start := (int(ch) - 32) * 9
			fmt.Print(line[start+row])
		}
	}
	fmt.Println()
}
