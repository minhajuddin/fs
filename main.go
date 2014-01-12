package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Simple static webserver:
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	log.Println("Starting server at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir(pwd))))
}
