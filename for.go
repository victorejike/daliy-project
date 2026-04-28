package main

import (
	"fmt"
)

func main() {
	word := [3]string{"hello", "world!", "victor"}
	for _, i := range word {
		fmt.Printf("%v \n", len(i))
	}

}
