package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

const url = "https://e-tutor-wine.vercel.app/"

func main() {
	fmt.Println("Welcome to web request in go lang")
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	fmt.Println(response.Status)
	fmt.Println("Response is  a type of ", reflect.TypeOf(response))

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	// Check if the response was OK
	fmt.Println(string(databytes))

}
