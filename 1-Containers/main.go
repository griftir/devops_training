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
	http.ListenAndServe(":8080", r)
}
