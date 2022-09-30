package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	for i := 0; i < 1000; i += 10 {
		conn, err := net.Dial("tcp", "192.168.214.132:21")
		if err != nil {
			log.Fatalf("[!] net.Dial error at offset %d: %s\n", i, err)
		}

		bufio.NewReader(conn).ReadString('\n')

		buf := ""

		for n := 0; n <= i; n++ {
			buf += "A"
		}

		fmt.Printf("%d chars\n", i)

		raw := "USER %s\r\n"
		fmt.Fprintf(conn, raw, buf)
		bufio.NewReader(conn).ReadString('\n')

		raw = "PASS test\r\n"
		fmt.Fprint(conn, raw)
		bufio.NewReader(conn).ReadString('\n')

		if err := conn.Close(); err != nil {
			fmt.Printf("[*] conn.Close error at offset %d\n", i)
		}

	}
}
