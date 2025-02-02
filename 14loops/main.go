package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for _, day := range days {
		fmt.Printf("index is and value is %v\n", day)
	}

	rougueValue := 1

	for rougueValue < 10 {
		if rougueValue == 4 {
			goto lco
		}
		fmt.Println(rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("I will start a youtube channel.")
}
