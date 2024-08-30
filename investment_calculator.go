package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (futureValue float64, futureRealValue float64) {
	futureValue = investmentAmount * math.Pow(1+expectedReturnRate/float64(100), years)
	futureRealValue = futureValue / math.Pow(1+inflationRate/float64(100), years)
	return futureValue, futureRealValue
}

func getUserInput(requestedInput string) float64 {
	var userInput float64
	fmt.Printf("Enter %s: \n", requestedInput)
	fmt.Scan(&userInput)
	return userInput
}

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	investmentAmount = getUserInput("investment amount")
	expectedReturnRate = getUserInput("rate of return")
	years = getUserInput("investment length of years")

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	fmt.Printf("Future value: %.2f\n", futureValue)
	fmt.Printf("Future real value(adjusted for Inflation): %.2f\n", futureRealValue)
}
