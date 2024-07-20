package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount, years float64 = 1000, 10
	// var expectedReturnRatePerc = 5.5
	// var years float64 = 10

	// a convetion for declaring a variable with a type inferred by go
	expectedReturnRatePerc := 5.5

	futureValue := investmentAmount * math.Pow(1+expectedReturnRatePerc/100, years)
	fmt.Println(futureValue)
}
