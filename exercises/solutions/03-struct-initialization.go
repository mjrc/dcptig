package main

import "fmt"

func main() {

	type Greeter struct {
		Expression string
		Subject    string
	}

	hw := Greeter{
		Expression: "Hello",
		Subject:    "World",
	}

	fmt.Printf("%s %s\n", hw.Expression, hw.Subject)
}
