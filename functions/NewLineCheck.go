package handler

import "strings"

// NewLineCheck verifies if the input string contains only newline sequences.
func NewLineCheck(str string) bool {
	cont := 0
	// Replace all occurrences of "\r\n" with a special character
	result := strings.ReplaceAll(str, `\r\n`, "é")
	
	// Count occurrences of the special character
	for i := 0; i < len(result); i++ {
		if result[i] == 'é' {
			cont++
		}
	}
	
	// Return true if the count matches the length of the modified string
	return cont == len(result)
}
