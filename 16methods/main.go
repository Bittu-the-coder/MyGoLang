package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs in the go lang")
	//no inheritance in go lang : no super or parent
	bittu := User{"Bittu", "bittu@gmail.com", true, 18}
	fmt.Println(bittu)
	// //accessing fields
	// fmt.Println(bittu.Name)
	// fmt.Println(bittu.Email)
	// fmt.Println(bittu.Status)
	// fmt.Println(bittu.Age)
	bittu.GetStatus()
	bittu.NewEmail()
	fmt.Println(bittu)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active : ", u.Status)
}

func (u User) NewEmail() {
	u.Email = "test@ai.com"
	fmt.Println("The new email is : ", u.Email)
}
