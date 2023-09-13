package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", helloHandler)

	http.ListenAndServe(":8080", server)
}
