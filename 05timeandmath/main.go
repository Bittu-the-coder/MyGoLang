package main

import (
	"crypto"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(createDate.Format("01-02-2006 15:04:05 Monday"))

	//random number
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(5) + 1)

	//random from crypto/rand
	cryptoRand := crypto.Reader
	randomNumber := cryptoRand.Int63()
	fmt.Println(randomNumber)

	// mynum, _ := rand.Int(rand.Reader, big.NewInt(5))
	// fmt.Println(mynum.Int64())
}
