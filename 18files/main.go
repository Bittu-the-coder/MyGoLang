package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in go lang")

	content := "This need to go in a file - LearnCodeOnline.in"

	file, err := os.Create("./myTxtFile.txt")

	checkNilError(err)

	lenght, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("Lenght is : ", lenght)
	readFile("./myTxtFile.txt")
	file.Close()
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println("Text data inside is : ", string(databyte))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
