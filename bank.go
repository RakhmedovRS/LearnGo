package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFileName = "balance.txt"
const defaultBalance = 1000

func persistBalanceIntoFile(balance float64) {
	_ = os.WriteFile(balanceFileName, []byte(fmt.Sprintf("%f", balance)), 0644)
}

func readBalanceFromFile() (float64, error) {
	bytes, err := os.ReadFile(balanceFileName)
	if err != nil {
		return defaultBalance, errors.New("unable to find balance file")
	}
	balanceText := string(bytes)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return defaultBalance, errors.New("unable to parse balance file")
	}
	return balance, nil
}

func bank() {
	fmt.Println("Welcome to Go Bank!")
mainLoop:
	for {
		var balance, err = readBalanceFromFile()
		if err != nil {
			fmt.Println(err)
			panic(err)
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
			persistBalanceIntoFile(balance)
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
			persistBalanceIntoFile(balance)
			fmt.Printf("Your new balance: %f\n", balance)
		case 4:
			fmt.Println("Goodbye!")
			break mainLoop
		default:
			panic(errors.New("invalid input"))
		}
	}
}
