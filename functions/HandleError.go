package handler

import (
	"html/template"
	"net/http"
)

func handleError(w http.ResponseWriter, msg string, codeError int) {
	tpStatusError, er_status := template.ParseFiles("Static/error.html")
	w.WriteHeader(codeError) // for console
	if er_status != nil {
		http.Error(w, "Internal Server Error 500 :: File Not Found", http.StatusInternalServerError)
		return
	}
	err := tpStatusError.Execute(w, msg)
	if err != nil {
		handleError(w, "Internal Server Error 500", http.StatusInternalServerError)
	}
}
