package main

import (
	"fmt"
	"net"
)

func main() {
	// Waits for a timeout
	openports := []int{}
	fmt.Println("Starting port scan")
	for port := 1; port < 25; port++ {
		_, err := net.Dial("tcp", fmt.Sprintf("scanme.nmap.org:%d", port))
		fmt.Printf("port %d: %s\n", port, err)
		if err != nil {
			openports = append(openports, port)
		}
	}
	for _, port := range openports {
		fmt.Printf("%d open \n", port)
	}
}
