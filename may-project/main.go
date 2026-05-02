package main

import (
	"fmt"
)

func main() {
	count := 10

	for i := 0; i < count; i++ {
		if i != 20 {
			fmt.Println(count)
		}
	}
}
