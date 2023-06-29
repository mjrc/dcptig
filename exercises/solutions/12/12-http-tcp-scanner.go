package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sort"

	"github.com/gorilla/mux"
)

func worker(target string, ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func scan(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	target := vars["ip"]

	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(target, ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			fmt.Println(i)
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results

		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open \n", port)
	}

	fmt.Fprintf(w, "%v", openports)
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
	r.HandleFunc("/scan/{target}", scan).Methods("GET")

	fmt.Println("Listening on port 8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln("There's an error with the server", err)
	}
}
