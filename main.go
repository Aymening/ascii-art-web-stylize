package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	handler "handler/functions"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Error: go run .")
		return
	}

	http.Handle("/style.css", http.FileServer(http.Dir("Static")))
	http.Handle("/bg1.jpg", http.FileServer(http.Dir("Static")))

	http.HandleFunc("/", handler.AsciiArtHandler)
	http.HandleFunc("/ascii-art", handler.ResultHandler)
	fmt.Println("http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
