package handler

import (
	"strings"
)

// GeneratePattern creates an ASCII art pattern based on the input string and the provided banner template.
func GeneratePattern(input string, Banner string) string {
	// Split the banner template by double newlines
	SplitFile := strings.Split(Banner, "\n\n")
	
	// Split the input string by carriage return and newline
	SplitArg := strings.Split(input, "\r\n")
	
	// Initialize a 2D slice to store the result pattern
	result := make([][]string, len(SplitArg))
	for mak := 0; mak < len(SplitArg); mak++ {
		result[mak] = make([]string, 8)
	}
	
	s := ""
	
	// Iterate through each line of the input
	for i := 0; i < len(SplitArg); i++ {
		// Check if the line contains valid characters
		if Check(SplitArg[i]) != "" {
			return Check(SplitArg[i])
		}
		// Generate the ASCII art for each line
		for l := 0; l < 8; l++ {
			for j := 0; j < len(SplitArg[i]); j++ {
				SplitCharacters := strings.Split(string(SplitFile[SplitArg[i][j]-32]), "\n")
				if SplitCharacters[l] == "" {
					SplitCharacters[l] = "      "
				}
				result[i][l] += SplitCharacters[l]
			}
			if result[i][l] != "" {
				s = s + string(result[i][l]) + "\n"
			} else {
				s = s + "\n"
				break
			}
		}
	}
	
	// Return the final ASCII art pattern
	return s
}
