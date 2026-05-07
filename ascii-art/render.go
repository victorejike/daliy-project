// render.go
package main

func RenderLine(s string, banner map[rune][]string) []string {
	result := make([]string, 8)

	for _, r := range s {
		char := banner[r]

		for i := 0; i < 8; i++ {
			result[i] += char[i]
		}
	}

	return result
}
