// TODO: This solution needs to be fixed

package main

import (
	"fmt"
	"time"
)

func hello(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "Hello"
}

func world(c chan string) {
	time.Sleep(1 * time.Second)
	c <- "World"
}

func main() {

	c := make(chan string)

	go hello(c)
	go world(c)

	subject, expression := <-c, <-c

	close(c)
	fmt.Printf("%v % v\n", subject, expression)

}
