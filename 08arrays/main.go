package main

import "fmt"

func main(){
	fmt.Println("Welcome to array in go lang")

	var fruitList [4]string
  
	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[2] = "Grapes"

	fmt.Println("Fruit List is : ", fruitList)
	fmt.Println("Fruit List is : ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mashrum"}
  
	fmt.Println("Vegitable List: ", vegList)
	fmt.Println("Vegitable List: ", len(vegList))
}