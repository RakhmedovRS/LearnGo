package main

import (
	"errors"
	"fmt"
)

func bank() {

	var balance = 1000

	fmt.Println("Welcome to Go Bank!")
mainLoop:
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")
		fmt.Print("Your choice: ")

		var userInput int
		fmt.Scanf("%d", &userInput)
		switch userInput {
		case 1:
			fmt.Printf("Your balance: %d\n", balance)
		case 2:
			fmt.Printf("How much do you want to deposit?\n")
			var depositAmount int
			fmt.Scanf("%d", &depositAmount)
			if depositAmount <= 0 {
				fmt.Printf("Your depositAmount %d is incorrect\n", depositAmount)
				continue mainLoop
			}
			balance += depositAmount
			fmt.Printf("Your new balance: %d\n", balance)
		case 3:
			fmt.Printf("How much do you want to withdraw?\n")
			var withdrawAmount int
			fmt.Scanf("%d", &withdrawAmount)
			if withdrawAmount > balance {
				fmt.Printf("Your withdrawAmount %d is incorrect\n", withdrawAmount)
				continue mainLoop
			}
			balance -= withdrawAmount
			fmt.Printf("Your new balance: %d\n", balance)
		case 4:
			fmt.Println("Goodbye!")
			break mainLoop
		default:
			errors.New("Invalid input")
		}
	}
}
