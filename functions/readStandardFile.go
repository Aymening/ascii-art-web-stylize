package handler

import (
	"fmt"
	"os"
	"strings"
)

// ReadStandardFile reads the content of the specified banner file and returns it as a string.
func ReadStandardFile(Banner string) string {
	// Read the content of the banner file
	cont, err := os.ReadFile("Banner/" + Banner + ".txt")
	if err != nil {
		// Print error message and exit if reading the file fails
		fmt.Println("PROGRAM : ", err)
		os.Exit(0)
	}
	
	// Replace all carriage return characters with an empty string
	banner := strings.ReplaceAll(string(cont), "\r", "")
	
	// Return the processed banner content
	return string(banner)
}
