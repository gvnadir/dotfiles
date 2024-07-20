package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000
	var expectedReturnRatePerc = 5.5
	var years float64 = 10

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRatePerc/100, years)
	fmt.Println(futureValue)
}
