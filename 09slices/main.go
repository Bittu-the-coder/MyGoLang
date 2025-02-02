package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
    fmt.Println("Welcome to slices in Go lang")

    var fruitList []string = []string{"Apple", "Tomato", "Peach"}

    fmt.Println("Type of fruit List is  ", reflect.TypeOf(fruitList))

    fruitList = append(fruitList, "Banana", "Mango", "Grapes")
    fmt.Println("Updated List is ", fruitList)

    fruitList = fruitList[1:3]
    fmt.Println(fruitList)


		highScores := make([]int, 4)

		highScores[0] = 323
		highScores[1] = 456
		highScores[2] = 789
		highScores[3] = 123
		fmt.Println(highScores)

		fmt.Println(sort.IntsAreSorted(highScores))
		sort.Ints(highScores)
		fmt.Println(highScores)
		fmt.Println(sort.IntsAreSorted(highScores))
    
		// fmt.Println(sort.IsSorted(highScores))


		var courses = []string{"react js", "js", "swift", "python", "ruby"}

		fmt.Println(courses)

		var index int = 2

		courses = append(courses[:index], courses[index+1:]...)
		fmt.Println(courses)

}