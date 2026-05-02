package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("usage: go run main. [string] [banner]")
		return
	}

	input := os.Args[1]
	bannerName := "standard"

	if len(os.Args) != 3 {
		bannerName = 
	}
}
