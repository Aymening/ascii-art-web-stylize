package handler

import (
	"html/template"
	"net/http"
)

// AsciiArtHandler handles the HTTP request for the ASCII art web page.
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML template file
	tmp, err := template.ParseFiles("Static/ascii-art-web.html")
	if err != nil {
		handleError(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}

	// Check if the request URL path is "/"
	if r.URL.Path != "/" {
		handleError(w, "Oops we can't find that page 404", http.StatusNotFound)
		return
	}

	// Execute the template and send it to the response writer
	tmp.ExecuteTemplate(w, "ascii-art-web.html", nil)
}
