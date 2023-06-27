package main

import "fmt"

func world() string {
	return "World"
}

func main() {
	var foo = "Hello" 
	bar := world() 

	fmt.Println(foo + " " + bar)
	fmt.Printf("%s %s\n", foo, bar)
}