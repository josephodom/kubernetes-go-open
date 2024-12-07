package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `{"message":"Hello, World!"}`)
	})

	log.Fatalf("Error: %+v", http.ListenAndServe("0.0.0.0:8080", nil))
}
