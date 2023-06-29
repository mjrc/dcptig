package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	greetings := []string{
		"Hey",
		"Howdy",
		"Hello",
		"Wazzup",
		"Yo",
	}

	subjects := []string{
		"World",
		"People",
		"Friends",
		"Strangers",
	}

	capacity := cap(greetings)

	for _, value := range subjects {
		random := rand.Intn(capacity)
		greeting := greetings[random]

		if strings.HasPrefix(greeting, "H") {
			fmt.Printf("%v %s!\n", greetings[random], value)
		}
	}
}
