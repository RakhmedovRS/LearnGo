package main

import "fmt"

var applications = map[int]func(){
	1: investmentCalculator,
	2: bank,
}

func main() {
	fmt.Println("Which application do you want to run?")
	fmt.Println("1. Investment calculator")
	fmt.Println("2. Bank")
	var userInput int
	fmt.Scanf("%d", &userInput)
	application, ok := applications[userInput]
	if !ok {
		fmt.Printf("Your choice %d is not recognized", userInput)
		return
	}

	application()
}
