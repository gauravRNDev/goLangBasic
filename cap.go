package main

import (
	"fmt"
	"strings"
)

func capitalizeChar(n string) {
	splitedStrArr := strings.Split(n, " ")
	var capFullArr = []string{}
	for _, value := range splitedStrArr {
		firstChar := strings.ToUpper(value[0:1])
		restChar := value[1:]
		finalChar := firstChar + restChar
		capFullArr = append(capFullArr, finalChar)
	}
	result := strings.Join(capFullArr, " ")
	fmt.Println(result)
}

func findLargest(n []int64) {
	fmt.Println(n)
	currentLargest := n[0]

	for _, value := range n {
		if value > currentLargest {
			currentLargest = value
		}
	}
	println(currentLargest)
}
