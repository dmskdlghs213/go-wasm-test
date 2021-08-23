package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("listen on http://localhost:8080")
	http.Handle("/", http.FileServer(http.Dir("app")))
	http.ListenAndServe(":"+"8080", nil)
}
