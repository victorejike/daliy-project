package main

import (
	"fmt"
)

func main() {
	text := "hello world"

	for i := 0; i < len(text); i++ {
		fmt.Printf("byte at %d:  and %c\n", i, text[2])
	}

}
