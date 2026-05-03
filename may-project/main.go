package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "" {
		return
	}
	content, _ := os.ReadFile("standard.txt")
	banner := strings.Split(string(content), "\n")

	input := strings.ReplaceAll(os.Args[1], "\\n", "\n")
	parts := strings.Split(input, "\n")

	for _, word := range parts {
		if word == "" {
			fmt.Println()
			continue
		}
		for i := 0; i < 9; i++ {
			for _, r := range word {
				// Each char block in file is 9 lines (8 art + 1 empty)
				// Index = (char - 32) * 9 + 1 (to skip initial file newline)
				idx := int(r-32)*9 + 1 + i
				if idx < len(banner) && i < 8 {
					fmt.Print(banner[idx])
				} else if i == 8 {
					fmt.Print("") // The 9th line
				}
			}
			fmt.Println()
		}
	}
}
