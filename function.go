package main

import (
	"fmt"
)

func addition(x int, y int) (result int) {
	result = x % y
	return result
}

func main() {
	fmt.Println(addition(5, 10))
}
