package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switch case of go lang")
	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("Dice value is 2 and you can move 2 steps")
	case 3:
		fmt.Println("Dice value is 3 and you can move 3 steps")
		fallthrough
	case 4:
		fmt.Println("Dice value is 4 and you can move 4 steps")
		fallthrough
	case 5:
		fmt.Println("Dice value is 5 and you can move 5 steps")
	case 6:
		fmt.Println("Dice value is 6 and you can move 6 steps")
	default:
		fmt.Println("Invalid dice value")
	}
}
