package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}

func main() {

	http.HandleFunc("/mick", handler)

	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
