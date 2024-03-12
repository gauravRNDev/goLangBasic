package main

import (
	"fmt"
	"strings"
)

func makeFirstCapital(n string) {
	var splitedStrArr = []string{}
	fmt.Println(splitedStrArr)

}

func main() {
	// fmt.Println("Hello Mr.")
	// greet := "Hello"

	// // strings
	// var firstName string = "john"
	// lastName := "wick"

	// fmt.Println(firstName, lastName)

	// // int and float
	// var numFirst int = 11333
	// var score float32 = 11.33
	// fmt.Println(numFirst, score)

	// // formatting strings
	// // fmt.Printf("%v %v %v how are you doing today?", greet, firstName, lastName)

	// // look := fmt.Sprintf("%v is old %v", lastName, 22.3)
	// // fmt.Print(look)

	// arrays
	// var lotteries [5]int = [5]int{3, 5, 6, 12, 34}

	// // updating
	// lotteries[0] = 44

	// // getting range
	// fixedRang := lotteries[1:3]
	// fmt.Println(lotteries, fixedRang)

	// slices
	// var scores = []int{2, 3, 4, 5}

	// // adding
	// scores = append(scores, 3, 4, 4)
	// fmt.Println(scores)

	//Strings methods

	// var testing string = "hello john wick"
	// fmt.Println(strings.Contains(testing, "john"))

	// var originalText string = "Hello World is Great"
	// fmt.Println(strings.ReplaceAll(originalText, "Hello", "Hi"))
	// fmt.Println(strings.Split(originalText, ""))

	// var names = []string{"john", "wick", "peter", "parker"}

	// looping
	// for x := 0; x < len(names); x++ {
	// 	fmt.Println(names[x])
	// }

	// for index, value := range names {
	// 	if index == 1 {
	// 		fmt.Printf("this is index one")
	// 		fmt.Println(value)
	// 	}
	// 	fmt.Printf(value)
	// }

	currentName := "peter parker"
	var firstHalf = currentName[1:strings.Index(currentName, " ")]
	var firstLetterOfFirstHalf = strings.ToUpper(currentName[0:1])
	var finalFirstHalf string = fmt.Sprintf("%s%s", firstLetterOfFirstHalf, firstHalf)
	fmt.Println(finalFirstHalf)

	capitalizeChar("peter parker")
	capitalizeChar("john wick killer")

}
