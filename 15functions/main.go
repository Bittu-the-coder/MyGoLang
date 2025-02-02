package main

import "fmt"

func main() {
	fmt.Println("Welcome to function in golang")
	greeter("Bittu")

	result, msg := adder(23, 232)
	fmt.Println(result, msg)

	fmt.Println(proAdder(23, 2343, 6, 7564, 424, 643, 43))
}

func proAdder(values ...int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum

}

func adder(val1 int, val2 int) (int, string) {
	return (val1 + val2), "Sum of two numbers"
}

func greeter(user string) {
	fmt.Print("Hello , ", user)
}
