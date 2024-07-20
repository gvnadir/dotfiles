package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	//var expectedReturnRatePerc = 5.5
	// a convetion for declaring a variable with a type inferred by go
	expectedReturnRatePerc := 5.5

	var years float64

	fmt.Print("Investment Amount: ")
	// the & symbol define a pointer and it's required here
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRatePerc)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRatePerc/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
