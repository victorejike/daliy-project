package main

import "fmt"

func ValidateInput(s string) (rune, error) {
	for _, r := range s {
		if r < 32 || r > 126 {
			return r, fmt.Errorf("unsupported character: %c", r)
		}
	}

	return 0, nil
}
