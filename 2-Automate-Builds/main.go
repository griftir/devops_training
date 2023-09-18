package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler)
	//print that the server is running
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
