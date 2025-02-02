package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now().Nanosecond()
	fmt.Println(presentTime)
	var count int = 5
	for i := 0; i < count; i++ {
				fmt.Println("*")
	}
}