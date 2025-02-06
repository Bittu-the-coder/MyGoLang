package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("Hello")

	myDefer()
}

// Hello two one World

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i, " ")
	}
}

// Hello 4 3 2 1 0 two one World
