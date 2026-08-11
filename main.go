package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go!")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/health", healthHandler)

	log.Println("Server running on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
