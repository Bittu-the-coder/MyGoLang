package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func main(){
  welcome := "Welcome to user input"
	var name string = "John"
	fmt.Println(welcome, name)
  
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza : ")

	// comma ok // err, err

	input, _ := reader.ReadString('\n')
	fmt.Println("You entered: ", input)
	fmt.Println("type of this rating is : ", reflect.TypeOf(input))
}