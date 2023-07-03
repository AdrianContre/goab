package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleDefault)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleDefault(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request:")
	fmt.Println("URL:", r.URL.String())
	fmt.Println("Method:", r.Method)
}
