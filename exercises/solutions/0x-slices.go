package main

import "fmt"

func main() {
	slice := []int{5, 7, 8}
	fmt.Println(slice)

	slice = append(slice, 1)
	fmt.Println(slice)

	slice = append(slice, 2)
	fmt.Println(slice)

	slice = append(slice, 3)
	fmt.Println(slice)

	slice = append(slice, 4)
	fmt.Println(slice)
}
