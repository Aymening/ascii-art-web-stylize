package handler

// HandleInput processes the input string and generates the corresponding ASCII art pattern using the provided banner template.
func HandleInput(arg string, Banner string) string {
	// Check if the input string contains newline characters
	if NewLineCheck(arg) {
		result := ""
		// Add carriage return and newline for each set of 4 characters in the input
		for i := 0; i < len(arg)/4; i++ {
			result += "\r\n"
		}
		return result
	}
	
	// Read the standard banner file content
	cont := ReadStandardFile(Banner)
	
	// Generate the ASCII art pattern
	return GeneratePattern(arg, cont)
}
