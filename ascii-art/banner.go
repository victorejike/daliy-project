// banner.go
package main

import (
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")

	banner := make(map[rune][]string)

	ascii := 32

	for i := 1; ascii <= 126 && i+7 < len(lines); i += 9 {
		banner[rune(ascii)] = lines[i : i+8]
		ascii++
	}

	return banner, nil
}
