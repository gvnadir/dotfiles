package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64 = 1000
	var years float64 = 10
	//var expectedReturnRatePerc = 5.5

	// a convetion for declaring a variable with a type inferred by go
	expectedReturnRatePerc := 5.5

	futureValue := investmentAmount * math.Pow(1+expectedReturnRatePerc/100, years)
	futureRealValue := futureValue * math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
