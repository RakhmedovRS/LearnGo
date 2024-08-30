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

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64
	fmt.Println("Enter investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Println("Enter rate of return: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Println("Enter length of years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	fmt.Printf("Future value: %.2f\n", futureValue)
	fmt.Printf("Future real value(adjusted for Inflation): %.2f\n", futureRealValue)
}
