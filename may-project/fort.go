package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Print("Enter your name: ")
	fmt.Scan(&input)

	var index int
	fmt.Print("Choose index: ")
	fmt.Scan(&index)

	switch index {
	case 0:
		fmt.Println(string(input[0]))
	case 1:
		fmt.Println(string(input[1]))
	}

}
