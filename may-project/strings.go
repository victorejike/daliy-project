package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "hello! victor how are you"

	group := strings.Split(text, "")

	fmt.Println(group[10])

	for index, char := range text {
		fmt.Printf("index : %d  character : %c\n", index, char)
	}
}
