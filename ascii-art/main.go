package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]

	_, err := ValidateInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := SplitInput(input)

	banner, err := LoadBanner("standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, line := range lines {
		if line == "" {
			fmt.Println()
			continue
		}

		RenderLine(line, banner)
	}
}
