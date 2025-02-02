package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps of go lang")
  language := make(map[string]string)

	language["js"] = "javascript"
	language["rb"] = "ruby"
	language["py"] = "python"
	language["go"] = "golang"

	fmt.Println("List of languages: ", language)
	fmt.Println("Js shorts for: ", language["js"])

	// loops are very interesting in go lang
	for _, value := range language{
		fmt.Println("shorts for: ", value)
	}
}