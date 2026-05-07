package main

import "strings"

func GenerateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}
	// split the input into lines using the SplitInput function

	parts := SplitInput(input)

	var result strings.Builder

	for i, part := range parts {

	
		if len(parts) == 2 && parts[0] == "" && parts[1] == "" {
			return "\n"
		}

	
		if part == "" {

			if i == len(parts)-1 {
				for j := 0; j < 8; j++ {
					result.WriteString("\n")
				}
				continue
			}

			
			result.WriteString("\n")
			continue
		}

		rendered := RenderLine(part, banner)

		for _, line := range rendered {
			result.WriteString(line)
			result.WriteString("\n")
		}
	}

	return result.String()
}
