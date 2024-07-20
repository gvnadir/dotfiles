package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
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

	// futureValue := investmentAmount * math.Pow(1+expectedReturnRatePerc/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRatePerc, years)

	// fmt.Println("FutureValue: ", futureValue)
	// fmt.Println("FutureRealValue", futureRealValue)

	// format text with printf
	// fmt.Printf("FutureValue: %v\nFutureRealValue: %v", futureValue, futureRealValue)

	// print 0 decimal place
	// fmt.Printf("FutureValue: %0.f\nFutureRealValue: %0.f", futureValue, futureRealValue)

	// store string using Sprint
	formattedFV := fmt.Sprintf("FutureValue: %0.f\n", futureValue)
	formattedFRV := fmt.Sprintf("FutureRealValue: %0.f\n", futureRealValue)
	fmt.Print(formattedFV, formattedFRV)
}

func calculateFutureValues(investmentAmount, expectedReturnRatePerc, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRatePerc/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
