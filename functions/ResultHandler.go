package handler

import (
	"html/template"
	"net/http"
)

// ResultHandler handles the HTTP POST request for generating and displaying the ASCII art result.
func ResultHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the result HTML template file
	tmp1, err := template.ParseFiles("Static/result.html")
	if err != nil {
		handleError(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}

	// Check if the request method is POST
	if r.Method != http.MethodPost {
		handleError(w, "BAD Request 400", http.StatusBadRequest)
		return
	}

	// Get the input value from the form
	input := r.FormValue("input")
	if len(input) > 150 || input == "" {
		handleError(w, "BAD Request 400", http.StatusBadRequest)
		return
	}

	// Get the banner value from the form
	Banner := r.FormValue("Banner")
	if Banner == "standard" || Banner == "benso" || Banner == "thinkertoy" || Banner == "shadow" {
		// Execute the template with the generated ASCII art
		err := tmp1.ExecuteTemplate(w, "result.html", HandleInput(input, Banner))
		if err != nil {
			handleError(w, "Internal Server Error 500", http.StatusInternalServerError)
			return
		}
	} else {
		handleError(w, "BAD Request 400", http.StatusBadRequest)
		return
	}
}
