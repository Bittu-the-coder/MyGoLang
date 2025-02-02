package main

import (
	"fmt"
)

func main() {
	fmt.Println("If else in golang")

	loginCout := 23
	var result string
	if loginCout > 18 {
		result = "You are a regular user"
		} else {
			result = "You are not regular user"
			}
			fmt.Println(result)

			if num :=3; num <12 {
				fmt.Println("Number is less than 12")
			} else {
				fmt.Println("Not less than 12")
			}
	
}