package main

import (
	"com.github.rakhmedovrs/udemy-go-course/fileops"
	"errors"
	"fmt"
)

const balanceFileName = "balance.txt"
const defaultBalance = 1000

func bank() {
	fmt.Println("Welcome to Go Bank!")
mainLoop:
	for {
		var balance, err = fileops.ReadFloatFromFile(balanceFileName, defaultBalance)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")
		fmt.Print("Your choice: ")

		var userInput int
		_, _ = fmt.Scanf("%d", &userInput)
		switch userInput {
		case 1:
			fmt.Printf("Your balance: %f\n", balance)
		case 2:
			fmt.Printf("How much do you want to deposit?\n")
			var depositAmount float64
			_, _ = fmt.Scanf("%f", &depositAmount)
			if depositAmount <= 0 {
				fmt.Printf("Your depositAmount %f is incorrect\n", depositAmount)
				continue mainLoop
			}
			balance += depositAmount
			fileops.PersistFloatIntoFile(balance, balanceFileName)
			fmt.Printf("Your new balance: %f\n", balance)
		case 3:
			fmt.Printf("How much do you want to withdraw?\n")
			var withdrawAmount float64
			_, _ = fmt.Scanf("%f", &withdrawAmount)
			if withdrawAmount > balance {
				fmt.Printf("Your withdrawAmount %f is incorrect\n", withdrawAmount)
				continue mainLoop
			}
			balance -= withdrawAmount
			fileops.PersistFloatIntoFile(balance, balanceFileName)
			fmt.Printf("Your new balance: %f\n", balance)
		case 4:
			fmt.Println("Goodbye!")
			break mainLoop
		default:
			panic(errors.New("invalid input"))
		}
	}
}
