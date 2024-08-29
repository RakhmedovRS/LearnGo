package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64
	fmt.Println("Enter investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Println("Enter rate of return: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Println("Enter length of years: ")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/float64(100), years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/float64(100), years)
	fmt.Printf("Future value: %f", futureValue)
	fmt.Println(futureRealValue)
}
