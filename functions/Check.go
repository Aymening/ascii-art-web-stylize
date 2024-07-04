package handler

// Check validates if the given string contains only valid ASCII characters.
func Check(str string) string {
	// Iterate through each character in the string
	for i := 0; i < len(str); i++ {
		// Check if the character is outside the valid ASCII range
		if str[i] < 32 || str[i] > 126 {
			return "Invalid characters"
		}
	}
	// Return an empty string if all characters are valid
	return ""
}
