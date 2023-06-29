package main

import "fmt"

type Greeter struct {
	Expression string
	Subject string 
}

func (g Greeter) Greet() {
	fmt.Printf("%s %s\n", g.Expression, g.Subject)
}

func main() {

	hw := Greeter{
		Expression: "Hello",
		Subject: "World",
	}

	hw.Greet() 
}