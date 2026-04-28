package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n*10)
}

func main() {
	fmt.Println(factorial(19))
}
