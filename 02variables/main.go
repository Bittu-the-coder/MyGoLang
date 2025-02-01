package main

import "fmt"

const LoginToken string = "ajgdnkf"

func main(){
	var username string = "bittu"
	fmt.Println(username);
	fmt.Print("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Println("Variable is of type: ", isLoggedIn)

	var num int = 10
	fmt.Println(num)
	fmt.Println("Variable is of type: ", num)

	var smallFloat float64 = 32.3421234
	fmt.Println(smallFloat)
	fmt.Println("Variable is of type: ", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Println("Variable is of type: ", anotherVariable)

	//implicit type
	var website = "https://e-tutor-wine.vercel.app/"
	fmt.Println(website)

	// no var style
	numberOfUser := 100
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Println("Variable is of type : ", LoginToken)
}