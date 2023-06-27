package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	greeting := fmt.Sprintf("Hello " + vars["name"])
	fmt.Fprintf(w, greeting)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	r.Use(loggingMiddleware)
	r.HandleFunc("/hello", handler).Methods("GET")
	r.HandleFunc("/hello/{name}", handler).Methods("GET")

	fmt.Println("Listening on port 8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln("There's an error with the server", err)
	}
}
