package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	// var expectedReturnRatePerc = 5.5
	// instead of using "var", we can use the convetion below for declaring a variable with a type inferred by go
	expectedReturnRatePerc := 5.5

	var years float64

	fmt.Print("Investment Amount: ")
	// the & symbol define a pointer and it's required here
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate (%): ")
	fmt.Scan(&expectedReturnRatePerc)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRatePerc/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// fmt.Println("FutureValue: ", futureValue)
	// fmt.Println("FutureRealValue", futureRealValue)

	// format text with printf
	// fmt.Printf("FutureValue: %v\nFutureRealValue: %v", futureValue, futureRealValue)
	fmt.Printf("FutureValue: %0.f\nFutureRealValue: %0.f", futureValue, futureRealValue)
}
